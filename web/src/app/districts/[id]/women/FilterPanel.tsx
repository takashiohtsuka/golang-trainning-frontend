"use client";

import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import WomanFilterPanel from "@/components/conditions/WomanFilterPanel";
import { fetchDistrictWomanCount } from "@/api/woman";
import { type BloodType } from "@/components/conditions/BloodTypeFilter";
import { type AgeRange } from "@/components/conditions/AgeFilter";
import { useDebounce } from "@/hooks/useDebounce";

type Props = {
  districtId: string;
  onSearch?: (bloodTypes: string[], ageRanges: string[]) => void;
  initialBloodTypes?: string[];
  initialAgeRanges?: string[];
};

function parseBloodTypes(values: string[]): BloodType[] {
  const VALID: BloodType[] = ["A", "B", "O", "AB"];
  return values.filter((v): v is BloodType => VALID.includes(v as BloodType));
}

function parseAgeRanges(values: string[]): AgeRange[] {
  return values.flatMap((v) => {
    const [min, max] = v.split("-").map(Number);
    return !isNaN(min) && !isNaN(max) ? [{ min, max }] : [];
  });
}

export default function FilterPanel({ districtId, onSearch, initialBloodTypes = [], initialAgeRanges = [] }: Props) {
  const [selectedBloodTypes, setSelectedBloodTypes] = useState<BloodType[]>(() => parseBloodTypes(initialBloodTypes));
  const [selectedAgeRanges, setSelectedAgeRanges] = useState<AgeRange[]>(() => parseAgeRanges(initialAgeRanges));
  const ageRangeStrings = selectedAgeRanges.map((r) => `${r.min}-${r.max}`);

  useEffect(() => {
    setSelectedBloodTypes(parseBloodTypes(initialBloodTypes));
    setSelectedAgeRanges(parseAgeRanges(initialAgeRanges));
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [initialBloodTypes.join(","), initialAgeRanges.join(",")]);

  const debouncedBloodTypes = useDebounce(selectedBloodTypes, 500);
  const debouncedAgeRangeStrings = useDebounce(ageRangeStrings, 500);

  const sortedBloodTypes = [...debouncedBloodTypes].sort();
  const sortedAgeRangeStrings = [...debouncedAgeRangeStrings].sort();

  const { data: total } = useQuery({
    queryKey: ["districtWomanCount", districtId, sortedBloodTypes, sortedAgeRangeStrings],
    queryFn: () => fetchDistrictWomanCount(districtId, sortedBloodTypes, sortedAgeRangeStrings).then((d) => d.total),
    staleTime: 30 * 1000,
  });

  return (
    <WomanFilterPanel
      selectedBloodTypes={selectedBloodTypes}
      onBloodTypesChange={setSelectedBloodTypes}
      selectedAgeRanges={selectedAgeRanges}
      onAgeRangesChange={setSelectedAgeRanges}
      total={total ?? null}
      onSearch={() => onSearch?.(selectedBloodTypes, ageRangeStrings)}
    />
  );
}
