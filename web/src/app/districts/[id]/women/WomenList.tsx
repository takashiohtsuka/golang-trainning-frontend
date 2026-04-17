"use client";

import { useEffect, useState } from "react";
import { fetchDistrictWomen } from "@/api/woman";
import type { WomanListItem } from "@/interfaces/district/womanList";

type Props = {
  districtId: string;
};

export default function WomenList({ districtId }: Props) {
  const [women, setWomen] = useState<WomanListItem[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetchDistrictWomen(districtId)
      .then((data) => setWomen(data.women))
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, [districtId]);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error}</p>;
  if (women.length === 0) return <p>女性が見つかりませんでした。</p>;

  return (
    <ul>
      {women.map((w) => (
        <li key={w.id}>
          <strong>{w.name}</strong>
          {w.age != null && <span> / {w.age}歳</span>}
          {w.birthplace && <span> / {w.birthplace}</span>}
          {w.blood_type && <span> / {w.blood_type}型</span>}
        </li>
      ))}
    </ul>
  );
}
