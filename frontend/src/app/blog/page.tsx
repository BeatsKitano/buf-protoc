'use client';

import { useQuery } from '@tanstack/react-query';

export default function Page() {
  const { data, error, isLoading } = useQuery({
    queryKey: ['data'], // 查询键，用于缓存和重新获取
    queryFn: async () => {
      const res = await fetch('https://api.vercel.app/blog');
      return res.json();
    },
  });

  if (error) return <div>请求失败</div>;
  if (isLoading) return <div>加载中...</div>;

  return (
    <>
    <ul>
      {data.map((blog) => (
        <li key={blog.id}>{blog.title}</li>
      ))}
    </ul>
    </>
  )
}