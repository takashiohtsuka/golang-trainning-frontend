import WomenList from "./WomenList";

type Props = {
  params: Promise<{ id: string }>;
};

export default async function Page({ params }: Props) {
  const { id } = await params;

  return (
    <main>
      <h1>District {id} の女性一覧</h1>
      <WomenList districtId={id} />
    </main>
  );
}
