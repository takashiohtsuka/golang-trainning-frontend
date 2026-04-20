"use client";

import { useEffect, useState } from "react";
import { fetchDistrictWomen } from "@/api/woman";
import type { WomanListItem } from "@/interfaces/district/womanList";

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
  const [women, setWomen] = useState<WomanListItem[]>([]);
  const [total, setTotal] = useState<number>(0);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    setLoading(true);
    fetchDistrictWomen(districtId, bloodTypes, ageRanges, page)
      .then((data) => {
        setWomen(data.women);
        setTotal(data.total);
      })
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, [districtId, bloodTypes, ageRanges, page]);

  return { women, total, loading, error };
}
