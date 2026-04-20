import WomenPageClient from "./WomenPageClient";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function Page({ params }: Props) {
  const { id } = await params;

  return (
    <main>
      <h1>District {id} の女性一覧</h1>
      <WomenPageClient districtId={id} />
    </main>
  );
}