export default function DashboardLayout({
    children,
  }: {
    children: React.ReactNode
  }) {
    return (
        <section>
            <p>blog布局模版</p>
            {children}
        </section>
    )
  }