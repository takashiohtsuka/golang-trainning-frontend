"use client";

import { useState, useEffect } from "react";
import { useQuery } from "@tanstack/react-query";
import FilterPanelShell from "@/components/conditions/FilterPanelShell";
import BloodTypeFilter, { type BloodType, parseBloodTypes } from "@/components/conditions/BloodTypeFilter";
import AgeFilter, { type AgeRange, parseAgeRanges } from "@/components/conditions/AgeFilter";
import { fetchDistrictWomanCount } from "@/api/woman";
import { useDebounce } from "@/hooks/useDebounce";

type Props = {
  districtId: string;
  onSearch?: (bloodTypes: string[], ageRanges: string[]) => void;
  initialBloodTypes?: string[];
  initialAgeRanges?: string[];
};

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
    <FilterPanelShell onSearch={() => onSearch?.(selectedBloodTypes, ageRangeStrings)}>
      <BloodTypeFilter selected={selectedBloodTypes} onChange={setSelectedBloodTypes} />
      <AgeFilter selected={selectedAgeRanges} onChange={setSelectedAgeRanges} />
      {total != null && <p>該当件数: {total}件</p>}
    </FilterPanelShell>
  );
}
