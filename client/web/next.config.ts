import type { NextConfig } from 'next';

const nextConfig: NextConfig = {
  /* config options here */
  images: {
    domains: ['image.tmdb.org'], // Thêm domain của ảnh
    unoptimized: false,
  },
};

export default nextConfig;
