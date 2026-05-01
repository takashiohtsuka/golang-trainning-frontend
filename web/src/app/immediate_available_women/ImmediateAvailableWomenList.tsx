"use client";

import { useImmediateAvailableWomen } from "./useImmediateAvailableWomen";

const PER_PAGE = 10;

type Props = {
  page: number;
  prefectureId?: number;
  districtId?: number;
  businessTypes: string[];
  bloodTypes: string[];
  ageRanges: string[];
  onPageChange: (page: number) => void;
};

export default function ImmediateAvailableWomenList({
  page,
  prefectureId,
  districtId,
  businessTypes,
  bloodTypes,
  ageRanges,
  onPageChange,
}: Props) {
  const { women, total, loading, error } = useImmediateAvailableWomen({
    page,
    prefectureId,
    districtId,
    businessTypes,
    bloodTypes,
    ageRanges,
  });

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (women.length === 0) return <p>該当する女性が見つかりませんでした。</p>;

  const totalPages = Math.ceil(total / PER_PAGE);

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
            <span> / {w.store.name}（{w.store.business_type}）</span>
            {w.expires_at && <span> / 受付終了: {new Date(w.expires_at).toLocaleTimeString("ja-JP")}</span>}
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
