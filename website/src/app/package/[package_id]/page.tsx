import { Suspense } from "react"
import { notFound } from "next/navigation"
import { Package } from "../../page"

type Variation = {
    ID: number
    PackageID: number
    DistroID: number
    Name: string
    Version: string
    PackageUrl: string
    DownloadUrl: string
    DistroName: string
    DistroVersion: string
}

async function getPackageInfo(package_id: string): Promise<Package> {
    console.log(`${process.env.API_URL}/packages/${package_id}`)
    const res = await fetch(`${process.env.API_URL}/packages/${package_id}`,
        {
            headers: {
                'Cache-Control': 'no-cache'
            }
        });

    if (!res.ok) {
        console.log(`error querying '${process.env.API_URL}/packages/${package_id}'`)
        throw new Error('Failed to fetch package info');
    }

    const data = await res.json();

    return data;
}

async function PackageInfo({ data }: { data: Package }) {
    return (
        <>
            <h1>{data.HumanName}</h1>
            <p>{data.Description}</p>
        </>
    );
}

async function getVariations(package_id: string): Promise<Variation[]> {
    const res = await fetch(`${process.env.API_URL}/packages/${package_id}/variations`,
        {
            headers: {
                'Cache-Control': 'no-cache'
            }
        });

    if (!res.ok) {
        console.log(`error querying '${process.env.API_URL}/packages/${package_id}/variations'`)
        throw new Error('Failed to fetch variations');
    }

    const data = await res.json();

    if (data && typeof data === 'object' && !Array.isArray(data)) {
        return Object.values(data);
    }

    return data;
}

async function Variations({ data }: { data: Variation[] }) {
    if (!Array.isArray(data)) {
        return <p>No variations available</p>;
    }

    return (
        <>
            {data.map((variation: Variation) => (
                <div key={variation.ID} style={{"padding": "10px 0px"}}>
                    <h1>{variation.DistroName} {variation.DistroVersion}</h1>
                    <p>{variation.Name}</p>
                </div>
            ))}
        </>
    );
}

export default async function Page({ params }: { params: Promise<{ package_id: string }> }) {
    const { package_id } = await params;

    // Validate that package_id is numeric
    if (!/^\d+$/.test(package_id)) {
        notFound();
    }

    const packageInfo = await getPackageInfo(package_id);
    const variations = await getVariations(package_id);

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