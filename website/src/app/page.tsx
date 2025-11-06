import Link from "next/link"
import { Suspense } from "react"

export type Package = {
    ID: number
    HumanName: string
    Name: string
    LatestVersion: string
    Description: string
    Keywords: string[]
    Homepage: string
    Developer: string[]
}

async function getPackages() {
    const res = await fetch(`${process.env.API_URL}/packages/list`,
        {
            headers: {
                'Cache-Control': 'no-cache'
            }
        })

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

async function Packages({ data }: { data: Package[] }) {
    return (
        <>
            {data.map((pkg: Package) => (
                <Link href={`/package/${pkg.ID}`} key={pkg.ID} className="m-8 block px-6 py-4 border-gray-200 rounded-lg shadow border hover:bg-gray-900 dark:bg-gray-900 dark:border-gray-800 dark:hover:bg-gray-800 w-96">
                    <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{pkg.HumanName}</h5>
                    <p className="font-normal text-gray-700 dark:text-gray-400">{pkg.Description}</p>
                </Link>
            ))
            }
        </>
    )
}

export default async function Home() {

    const data = await getPackages()

    return (
        <main className="min-h-screen">
            <Suspense fallback={<div>Loading...</div>}>
                <div className="m-auto flex flex-row flex-wrap items-stretch">
                    <Packages data={data} />
                </div>
            </Suspense>
        </main >
    )
}


