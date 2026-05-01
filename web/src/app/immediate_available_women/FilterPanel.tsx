"use client";

import { useState, useEffect } from "react";
import FilterPanelShell from "@/components/conditions/FilterPanelShell";
import PrefectureSelect from "@/components/conditions/PrefectureSelect";
import DistrictSelect from "@/components/conditions/DistrictSelect";
import BusinessTypeFilter from "@/components/conditions/BusinessTypeFilter";
import BloodTypeFilter, { type BloodType, parseBloodTypes } from "@/components/conditions/BloodTypeFilter";
import AgeFilter, { type AgeRange, parseAgeRanges } from "@/components/conditions/AgeFilter";
import type { PrefectureItem } from "@/interfaces/prefecture/list";
import type { BusinessTypeItem } from "@/interfaces/businessType/list";

type Props = {
  prefectures: PrefectureItem[];
  businessTypes: BusinessTypeItem[];
  onSearch: (params: {
    prefectureId?: number;
    districtId?: number;
    businessTypes: string[];
    bloodTypes: string[];
    ageRanges: string[];
  }) => void;
  initialPrefectureId?: number;
  initialDistrictId?: number;
  initialBusinessTypes?: string[];
  initialBloodTypes?: string[];
  initialAgeRanges?: string[];
};

export default function FilterPanel({
  prefectures,
  businessTypes,
  onSearch,
  initialPrefectureId,
  initialDistrictId,
  initialBusinessTypes = [],
  initialBloodTypes = [],
  initialAgeRanges = [],
}: Props) {
  const [selectedPrefectureId, setSelectedPrefectureId] = useState<number | undefined>(initialPrefectureId);
  const [selectedDistrictId, setSelectedDistrictId] = useState<number | undefined>(initialDistrictId);
  const [selectedBusinessTypes, setSelectedBusinessTypes] = useState<string[]>(initialBusinessTypes);
  const [selectedBloodTypes, setSelectedBloodTypes] = useState<BloodType[]>(() => parseBloodTypes(initialBloodTypes));
  const [selectedAgeRanges, setSelectedAgeRanges] = useState<AgeRange[]>(() => parseAgeRanges(initialAgeRanges));

  useEffect(() => {
    setSelectedDistrictId(undefined);
  }, [selectedPrefectureId]);

  const handleSearch = () => {
    onSearch({
      prefectureId: selectedPrefectureId,
      districtId: selectedDistrictId,
      businessTypes: selectedBusinessTypes,
      bloodTypes: selectedBloodTypes,
      ageRanges: selectedAgeRanges.map((r) => `${r.min}-${r.max}`),
    });
  };

  return (
    <FilterPanelShell onSearch={handleSearch}>
      <PrefectureSelect prefectures={prefectures} value={selectedPrefectureId} onChange={setSelectedPrefectureId} />
      <DistrictSelect prefectureId={selectedPrefectureId} value={selectedDistrictId} onChange={setSelectedDistrictId} />
      <BusinessTypeFilter businessTypes={businessTypes} selected={selectedBusinessTypes} onChange={setSelectedBusinessTypes} />
      <BloodTypeFilter selected={selectedBloodTypes} onChange={setSelectedBloodTypes} />
      <AgeFilter selected={selectedAgeRanges} onChange={setSelectedAgeRanges} />
    </FilterPanelShell>
  );
}
