#!/usr/bin/env python3
"""
Parse go test -bench outputs and generate Markdown tables and a compatibility matrix.

Inputs:
    - benchmark_threeLevelsWithIndex_test.result.txt
    - benchmark_recursiveDescentWithFilter_test.result.txt

Outputs (written to assets/):
    - bench_table_simple.md
    - bench_table_complex.md
    - support_matrix.md
    - bench_chart_simple.svg
    - bench_chart_complex.svg

"""

from __future__ import annotations

import re
from dataclasses import dataclass
from pathlib import Path
from typing import Dict, List, Tuple
from tabulate import tabulate
import pygal
from pygal.style import Style


ROOT = Path(__file__).resolve().parents[1]
ASSETS = ROOT / "assets"


@dataclass
class Bench:
    name: str
    iters: int
    ns_per_op: float
    b_per_op: int
    allocs_per_op: int


LIBRARY_NAME_PAIRS: List[Tuple[str, str]] = [
    ("AsaiYusuke_JSONPath_reuseBuffer", "AsaiYusuke/JSONPath (reuse)"),
    ("AsaiYusuke_JSONPath", "AsaiYusuke/JSONPath"),
    ("PaesslerAG_JSONPath", "PaesslerAG/JSONPath"),
    ("bhmj_JSON_Slice", "bhmj/JSONSlice"),
    ("ohler55_OjG_jp", "ohler55/OjG (jp)"),
    ("oliveagle_JsonPath", "oliveagle/JsonPath"),
    ("Spyzhov_Abstract_JSON", "Spyzhov/ajson"),
    ("vmware_labs_YAML_JSONPath", "vmware-labs/YAML JSONPath"),
]
DISPLAY_NAME = {k: v for k, v in LIBRARY_NAME_PAIRS}


LINE_RE = re.compile(
    r"^Benchmark\d+_(.+?)-\d+\s+(\d+)\s+([\d\.]+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op$"
)


def parse_result(path: Path) -> Dict[str, Bench]:
    benches: Dict[str, Bench] = {}
    if not path.exists():
        return benches
    for line in path.read_text().splitlines():
        m = LINE_RE.match(line.strip())
        if not m:
            continue
        name = m.group(1)
        iters = int(m.group(2))
        ns_per_op = float(m.group(3))
        b_per_op = int(m.group(4))
        allocs_per_op = int(m.group(5))
        benches[name] = Bench(name, iters, ns_per_op, b_per_op, allocs_per_op)
    return benches


def render_perf_table(benches: Dict[str, Bench], out_path: Path) -> None:
    if not benches:
        out_path.write_text("_No results found._\n")
        return

    sorted_items = sorted(benches.values(), key=lambda b: b.ns_per_op)
    fastest = sorted_items[0].ns_per_op

    header = [
        "Rank",
        "Library",
        "Time (ns/op)",
        "Memory (B/op)",
        "Allocations (allocs/op)",
        "Relative speed (fastest = 1x)",
    ]
    rows: List[List[str]] = []
    for i, b in enumerate(sorted_items, start=1):
        disp = DISPLAY_NAME.get(b.name, b.name)
        rel = b.ns_per_op / fastest if fastest > 0 else 0.0
        rows.append(
            [
                str(i),
                disp,
                f"{b.ns_per_op:.2f}",
                str(b.b_per_op),
                str(b.allocs_per_op),
                f"{rel:.2f}x",
            ]
        )

    colalign = ("center", "left", "right", "right", "right", "right")
    table_md = tabulate(
        rows,
        headers=header,
        tablefmt="pipe",
        colalign=colalign,
        disable_numparse=True,
    )
    out_path.write_text(table_md)


def render_support_matrix(
    simple: Dict[str, Bench], complex_: Dict[str, Bench], out_path: Path
) -> None:
    header = ["Library", "Simple query", "Complex query"]
    rows: List[List[str]] = []
    for key, disp in LIBRARY_NAME_PAIRS:
        ok_simple = "✅" if key in simple else "❌"
        ok_complex = "✅" if key in complex_ else "❌"
        rows.append([disp, ok_simple, ok_complex])
    table_md = tabulate(rows, headers=header, tablefmt="pipe")
    out_path.write_text(table_md)


_PALETTE = [
    "#4e79a7",
    "#f28e2b",
    "#e15759",
    "#76b7b2",
    "#59a14f",
    "#edc948",
    "#b07aa1",
    "#ff9da7",
    "#9c755f",
    "#bab0ab",
]

COLOR_MAP = {
    key: _PALETTE[i % len(_PALETTE)]
    for i, (key, _disp) in enumerate(LIBRARY_NAME_PAIRS)
}


def render_svg_bar_chart(benches: Dict[str, Bench], out_path: Path) -> None:
    if not benches:
        out_path.write_text(
            _svg_wrap(
                420,
                80,
                [
                    '<text x="20" y="50" font-family="sans-serif" font-size="14">No results found</text>'
                ],
            )
        )
        return

    items = sorted(benches.values(), key=lambda b: b.ns_per_op)
    items = list(reversed(items))
    labels = [DISPLAY_NAME.get(b.name, b.name) for b in items]
    values = [b.ns_per_op for b in items]

    n = len(items)
    width = 920
    base_px = 50
    slot_px = 25
    height = max(160, base_px + n * slot_px)

    chart = pygal.HorizontalBar(
        style=Style(
            label_font_size=14,
            value_font_size=14,
        ),
        show_legend=False,
        show_x_labels=False,
        show_y_labels=True,
        print_values=True,
        print_values_position="top",
        width=width,
        height=height,
        margin=0,
        spacing=8,
    )

    chart.x_labels = labels

    series_data = [
        {"value": val, "color": _color_for(item.name)}
        for val, item in zip(values, items)
    ]
    chart.add("", series_data)

    chart.render_to_file(str(out_path))


def _svg_wrap(width: int, height: int, elements: List[str]) -> str:
    bg = f'<rect x="0" y="0" width="{width}" height="{height}" fill="#ffffff" />'
    return (
        f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">'
        + bg
        + "".join(elements)
        + "</svg>"
    )


def _color_for(name: str) -> str:
    return COLOR_MAP.get(name, _PALETTE[0])


def main() -> None:
    simple_path = ROOT / "benchmark_threeLevelsWithIndex_test.result.txt"
    complex_path = ROOT / "benchmark_recursiveDescentWithFilter_test.result.txt"

    simple_result = parse_result(simple_path)
    complex_result = parse_result(complex_path)

    ASSETS.mkdir(exist_ok=True)

    render_perf_table(simple_result, ASSETS / "bench_table_simple.md")
    render_perf_table(complex_result, ASSETS / "bench_table_complex.md")

    render_support_matrix(simple_result, complex_result, ASSETS / "support_matrix.md")

    render_svg_bar_chart(simple_result, ASSETS / "bench_chart_simple.svg")
    render_svg_bar_chart(complex_result, ASSETS / "bench_chart_complex.svg")


if __name__ == "__main__":
    main()
