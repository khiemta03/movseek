import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  images: {
    domains: ['image.tmdb.org'], // Thêm domain của ảnh
    unoptimized: true,
  },
};

export default nextConfig;
