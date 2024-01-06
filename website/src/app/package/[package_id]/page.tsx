import { Suspense } from "react"
import { Package } from "../../page"
import Link from "next/link"

type Variation = {
    id: number
    package_id: number
    distro_id: number
    name: string
    version: string
    package_url: string
    download_url: string
    distro_name: string
    distro_version: string
}

async function getPackageInfo(package_id: number) {
    const res = await fetch(`${process.env.API_URL}/package/${package_id}`,
        {
            headers: {
                'Cache-Control': 'no-cache'
            }
        });

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data');
    }

    return res.json();
}

async function PackageInfo({ data }: { data: Package }) {
    return (
        <>
            <h1>{data.human_name}</h1>
            <p>{data.description}</p>
        </>
    );
}

async function getVariations(package_id: number) {
    const res = await fetch(`${process.env.API_URL}/package/${package_id}/variations`,
        {
            headers: {
                'Cache-Control': 'no-cache'
            }
        });

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json();
}

async function Variations({ data }: { data: Variation[] }) {
    return (
        <>
            {data.map((variation: Variation) => (
                <div key={variation.id}>
                    <h1>{variation.distro_name} {variation.distro_version}</h1>
                    <p>{variation.name}</p>
                </div>
            ))}
        </>
    );
}

export default async function Page({ params }: { params: { package_id: number } }) {

    let packageInfo = await getPackageInfo(params.package_id);
    packageInfo = packageInfo[0];
    const variations = await getVariations(params.package_id);

    return (
        <main>

            <Suspense fallback={<div>Loading...</div>}>
                <PackageInfo data={packageInfo} />
            </Suspense>

            <br /><br />

            <Suspense fallback={<div>Loading...</div>}>
                <Variations data={variations} />
            </Suspense>

        </main>
    );
}
