/** @type {import('next').NextConfig} */
const nextConfig = {
    async headers() {
        const cacheHeaders =
            process.env.NODE_ENV === "production"
                ? [{ key: "Cache-Control", value: "public, max-age=31536000, immutable" }] // Enable caching in production
                : [
                      { key: "Cache-Control", value: "no-store, no-cache, must-revalidate, max-age=0" },
                      { key: "Pragma", value: "no-cache" },
                      { key: "Expires", value: "0" },
                  ];

        return [
            {
                source: "/(.*)",
                headers: [
                    { key: "Access-Control-Allow-Credentials", value: "true" },
                    { key: "Access-Control-Allow-Origin", value: "http://localhost:*" },
                    ...cacheHeaders,
                ],
            },
        ];
    },
};

export default nextConfig;
