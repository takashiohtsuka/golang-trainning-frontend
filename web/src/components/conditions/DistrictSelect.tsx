"use client";

import { useQuery } from "@tanstack/react-query";
import { fetchDistricts } from "@/api/district";

type Props = {
  prefectureId: number | undefined;
  value: number | undefined;
  onChange: (value: number | undefined) => void;
};

export default function DistrictSelect({ prefectureId, value, onChange }: Props) {
  const { data } = useQuery({
    queryKey: ["districts", prefectureId],
    queryFn: () => fetchDistricts(prefectureId!),
    enabled: !!prefectureId,
    staleTime: 5 * 60 * 1000,
  });

  return (
    <div>
      <p>エリア：</p>
      <select
        value={value ?? ""}
        onChange={(e) => onChange(e.target.value ? Number(e.target.value) : undefined)}
        disabled={!prefectureId}
      >
        <option value="">指定なし</option>
        {data?.districts.map((d) => (
          <option key={d.id} value={d.id}>{d.name}</option>
        ))}
      </select>
    </div>
  );
}
