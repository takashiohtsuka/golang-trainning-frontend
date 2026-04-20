"use client";

import { useDistrictWomen } from "./useDistrictWomen";

type Props = {
  districtId: string;
  bloodTypes?: string[];
  ageRanges?: string[];
  page: number;
  onPageChange: (page: number) => void;
};

export default function WomenList({ districtId, bloodTypes = [], ageRanges = [], page, onPageChange }: Props) {
  const { women, total, loading, error } = useDistrictWomen({ districtId, bloodTypes, ageRanges, page });

  const perPage = 10;

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (women.length === 0) return <p>女性が見つかりませんでした。</p>;

  const totalPages = Math.ceil(total / perPage);

  return (
    <>
      <p>総件数: {total}件</p>
      <ul>
        {women.map((w) => (
          <li key={w.id}>
            <strong>{w.name}</strong>
            {w.age != null && <span> / {w.age}歳</span>}
            {w.birthplace && <span> / {w.birthplace}</span>}
            {w.blood_type && <span> / {w.blood_type}型</span>}
            {w.stores.length > 0 && (
              <span> / 店舗: {w.stores.map((s) => `${s.name}(${s.business_type})`).join(", ")}</span>
            )}
          </li>
        ))}
      </ul>
      <div>
        <button onClick={() => onPageChange(page - 1)} disabled={page <= 1}>前へ</button>
        <span> {page} / {totalPages} </span>
        <button onClick={() => onPageChange(page + 1)} disabled={page >= totalPages}>次へ</button>
      </div>
    </>
  );
}
