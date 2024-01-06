
export type PackagesResponse = {
    id: number
    human_name: string
    name: string
    latest_version: string
    description: string
    keywords: string[]
    homepage: string
    developer: string[]
}

async function getPackages() {
    const res = await fetch(`${process.env.API_URL}/packages`)

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default async function Home() {

    const data = await getPackages()

    console.log(data)

    return (
        <main>

            {data.map((pkg: PackagesResponse) => (
                <li key={pkg.id}>{pkg.human_name}</li>
            ))}

        </main>
    )
}
