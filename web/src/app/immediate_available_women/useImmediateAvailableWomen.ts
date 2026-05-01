"use client";

import { useQuery } from "@tanstack/react-query";
import { fetchImmediateAvailableWomen } from "@/api/immediateAvailableWoman";
import type { ImmediateAvailableWomanItem } from "@/interfaces/immediateAvailableWoman/list";

type Params = {
  page: number;
  prefectureId?: number;
  districtId?: number;
  businessTypes: string[];
  bloodTypes: string[];
  ageRanges: string[];
};

type ReturnValue = {
  women: ImmediateAvailableWomanItem[];
  total: number;
  loading: boolean;
  error: string | null;
};

export function useImmediateAvailableWomen(params: Params): ReturnValue {
  const { page, prefectureId, districtId, businessTypes, bloodTypes, ageRanges } = params;

  const sortedBusinessTypes = [...businessTypes].sort();
  const sortedBloodTypes = [...bloodTypes].sort();
  const sortedAgeRanges = [...ageRanges].sort();

  const { data, isLoading, error } = useQuery({
    queryKey: ["immediateAvailableWomen", page, prefectureId, districtId, sortedBusinessTypes, sortedBloodTypes, sortedAgeRanges],
    queryFn: () =>
      fetchImmediateAvailableWomen({
        page,
        prefectureId,
        districtId,
        businessTypes: sortedBusinessTypes,
        bloodTypes: sortedBloodTypes,
        ageRanges: sortedAgeRanges,
      }),
    staleTime: 30 * 1000,
  });

  return {
    women: data?.women ?? [],
    total: data?.total ?? 0,
    loading: isLoading,
    error: error?.message ?? null,
  };
}
