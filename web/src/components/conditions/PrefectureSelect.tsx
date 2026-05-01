"use client";

import type { PrefectureItem } from "@/interfaces/prefecture/list";

type Props = {
  prefectures: PrefectureItem[];
  value: number | undefined;
  onChange: (value: number | undefined) => void;
};

export default function PrefectureSelect({ prefectures, value, onChange }: Props) {
  return (
    <div>
      <p>都道府県：</p>
      <select
        value={value ?? ""}
        onChange={(e) => onChange(e.target.value ? Number(e.target.value) : undefined)}
      >
        <option value="">指定なし</option>
        {prefectures.map((p) => (
          <option key={p.id} value={p.id}>{p.name}</option>
        ))}
      </select>
    </div>
  );
}
