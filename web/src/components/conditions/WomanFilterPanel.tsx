"use client";

import { useState, useEffect } from "react";
import BloodTypeFilter, { type BloodType } from "@/components/conditions/BloodTypeFilter";
import AgeFilter, { type AgeRange } from "@/components/conditions/AgeFilter";

type FetchSearchCount = (bloodTypes: string[], ageRanges: string[]) => Promise<number>;

const VALID_BLOOD_TYPES: BloodType[] = ["A", "B", "O", "AB"];

function parseBloodTypes(values: string[]): BloodType[] {
  return values.filter((v): v is BloodType => VALID_BLOOD_TYPES.includes(v as BloodType));
}

function parseAgeRanges(values: string[]): AgeRange[] {
  return values.flatMap((v) => {
    const [min, max] = v.split("-").map(Number);
    return !isNaN(min) && !isNaN(max) ? [{ min, max }] : [];
  });
}

type Props = {
  fetchSearchCount: FetchSearchCount;
  onSearch?: (...args: Parameters<FetchSearchCount>) => void;
  initialBloodTypes?: string[];
  initialAgeRanges?: string[];
};

export default function FilterPanel({ fetchSearchCount, onSearch, initialBloodTypes = [], initialAgeRanges = [] }: Props) {
  const [selectedBloodTypes, setSelectedBloodTypes] = useState<BloodType[]>(() => parseBloodTypes(initialBloodTypes));
  const [selectedAgeRanges, setSelectedAgeRanges] = useState<AgeRange[]>(() => parseAgeRanges(initialAgeRanges));
  const [total, setTotal] = useState<number | null>(null);
  const ageRangeStrings = selectedAgeRanges.map((r) => `${r.min}-${r.max}`);

  useEffect(() => {
    fetchSearchCount(selectedBloodTypes, ageRangeStrings)
      .then((total) => setTotal(total))
      .catch(() => setTotal(null));
  }, [selectedBloodTypes, selectedAgeRanges, fetchSearchCount]);

  const search = () => {
    onSearch?.(selectedBloodTypes, ageRangeStrings);
  };

  return (
    <div>
      <BloodTypeFilter
        selected={selectedBloodTypes}
        onChange={setSelectedBloodTypes}
      />
      <AgeFilter
        selected={selectedAgeRanges}
        onChange={setSelectedAgeRanges}
      />
      {total !== null && <p>該当件数: {total}件</p>}
      <button onClick={search}>検索</button>
    </div>
  );
}
