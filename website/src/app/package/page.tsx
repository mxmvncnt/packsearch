import { PackagesResponse } from "../page"

async function getPackageInfo() {
    const res = await fetch(`${process.env.API_URL}/packages`)

    if (!res.ok) {
        // This will activate the closest `error.js` Error Boundary
        throw new Error('Failed to fetch data')
    }

    return res.json()
}

export default async function Package({ params }: { params: { package_id: number } }) {

    const data = await getPackageInfo()

    console.log(data)

    return (
        <main>

            {data.map((pkg: PackagesResponse) => (
                <li key={pkg.id}>{pkg.human_name}</li>
            ))}

        </main>
    )
}
