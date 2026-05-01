"use client";

import { useRouter, useSearchParams } from "next/navigation";
import { usePageParam } from "@/components/pagination/usePageParam";
import FilterPanel from "./FilterPanel";
import ImmediateAvailableWomenList from "./ImmediateAvailableWomenList";
import type { PrefectureItem } from "@/interfaces/prefecture/list";
import type { BusinessTypeItem } from "@/interfaces/businessType/list";

type Props = {
  prefectures: PrefectureItem[];
  businessTypes: BusinessTypeItem[];
};

export default function ImmediateAvailableWomenPageClient({ prefectures, businessTypes }: Props) {
  const router = useRouter();
  const searchParams = useSearchParams();
  const { page, goToPage } = usePageParam();

  const prefectureId = searchParams.get("prefecture_id") ? Number(searchParams.get("prefecture_id")) : undefined;
  const districtId = searchParams.get("district_id") ? Number(searchParams.get("district_id")) : undefined;
  const selectedBusinessTypes = searchParams.get("business_type")?.split(",").filter(Boolean) ?? [];
  const bloodTypes = searchParams.get("blood_type")?.split(",").filter(Boolean) ?? [];
  const ageRanges = searchParams.get("age_range")?.split(",").filter(Boolean) ?? [];

  const onSearch = (params: {
    prefectureId?: number;
    districtId?: number;
    businessTypes: string[];
    bloodTypes: string[];
    ageRanges: string[];
  }) => {
    const query = new URLSearchParams();
    query.set("page", "1");
    if (params.prefectureId) query.set("prefecture_id", String(params.prefectureId));
    if (params.districtId) query.set("district_id", String(params.districtId));
    if (params.businessTypes.length > 0) query.set("business_type", params.businessTypes.join(","));
    if (params.bloodTypes.length > 0) query.set("blood_type", params.bloodTypes.join(","));
    if (params.ageRanges.length > 0) query.set("age_range", params.ageRanges.join(","));
    router.push(`?${query.toString()}`);
  };

  return (
    <>
      <FilterPanel
        prefectures={prefectures}
        businessTypes={businessTypes}
        onSearch={onSearch}
        initialPrefectureId={prefectureId}
        initialDistrictId={districtId}
        initialBusinessTypes={selectedBusinessTypes}
        initialBloodTypes={bloodTypes}
        initialAgeRanges={ageRanges}
      />
      <ImmediateAvailableWomenList
        page={page}
        prefectureId={prefectureId}
        districtId={districtId}
        businessTypes={selectedBusinessTypes}
        bloodTypes={bloodTypes}
        ageRanges={ageRanges}
        onPageChange={goToPage}
      />
    </>
  );
}
