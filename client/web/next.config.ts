import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  images: {
    domains: ['image.tmdb.org', 'img.clerk.com'], // Thêm domain của ảnh
    unoptimized: false,
  },
  async rewrites() {
    return [
      {
        source: '/retriever/:path*', // Định nghĩa route mà bạn muốn proxy
        destination: 'https://awd-llm.azurewebsites.net/retriever/:path*', // Địa chỉ API backend
      },
    ];
  },
};

export default nextConfig;
