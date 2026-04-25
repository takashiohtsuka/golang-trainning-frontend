"use client";

import BloodTypeFilter, { type BloodType } from "@/components/conditions/BloodTypeFilter";
import AgeFilter, { type AgeRange } from "@/components/conditions/AgeFilter";

type Props = {
  selectedBloodTypes: BloodType[];
  onBloodTypesChange: (v: BloodType[]) => void;
  selectedAgeRanges: AgeRange[];
  onAgeRangesChange: (v: AgeRange[]) => void;
  total: number | null;
  onSearch: () => void;
};

export default function FilterPanel({ selectedBloodTypes, onBloodTypesChange, selectedAgeRanges, onAgeRangesChange, total, onSearch }: Props) {
  return (
    <div>
      <BloodTypeFilter
        selected={selectedBloodTypes}
        onChange={onBloodTypesChange}
      />
      <AgeFilter
        selected={selectedAgeRanges}
        onChange={onAgeRangesChange}
      />
      {total !== null && <p>該当件数: {total}件</p>}
      <button onClick={onSearch}>検索</button>
    </div>
  );
}
