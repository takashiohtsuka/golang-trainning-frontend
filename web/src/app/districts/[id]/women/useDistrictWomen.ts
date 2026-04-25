"use client";

import { useQuery } from "@tanstack/react-query";
import { fetchDistrictWomen } from "@/api/woman";
import type { WomanListItem, DistrictWomenResponse } from "@/interfaces/district/womanList";

type UseDistrictWomenParams = {
  districtId: string;
  bloodTypes: string[];
  ageRanges: string[];
  page: number;
};

type UseDistrictWomenReturnValue = {
  women: WomanListItem[];
  total: number;
  loading: boolean;
  error: string | null;
};

export function useDistrictWomen({ districtId, bloodTypes, ageRanges, page }: UseDistrictWomenParams): UseDistrictWomenReturnValue {
  const sortedBloodTypes = [...bloodTypes].sort();
  const sortedAgeRanges = [...ageRanges].sort();

  const { data, isLoading, error } = useQuery({
    queryKey: ["districtWomen", districtId, sortedBloodTypes, sortedAgeRanges, page],
    queryFn: () => fetchDistrictWomen(districtId, sortedBloodTypes, sortedAgeRanges, page),
    staleTime: 30 * 1000,
  });

  return {
    women: data?.women ?? [],
    total: data?.total ?? 0,
    loading: isLoading,
    error: error?.message ?? null,
  };
}
