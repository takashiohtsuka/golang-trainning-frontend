"use client";

type Props = {
  children: React.ReactNode;
  onSearch: () => void;
};

export default function FilterPanelShell({ children, onSearch }: Props) {
  return (
    <div>
      {children}
      <button onClick={onSearch}>検索</button>
    </div>
  );
}
