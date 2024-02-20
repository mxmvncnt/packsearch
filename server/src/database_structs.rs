use serde::Serialize;
use sqlx::{FromRow, Pool, Postgres};

pub struct AppState {
    pub(crate) db: Pool<Postgres>,
}

#[derive(Serialize, FromRow)]
pub struct Package {
    #[serde(skip_serializing)]
    id: i64,
    human_name: String,
    name: Option<String>,
    latest_version: Option<String>,
    description: Option<String>,
    keywords: Option<Vec<String>>,
    homepage: Option<String>,
    developer: Vec<String>,
}

#[derive(Serialize, FromRow)]
pub struct Distro {
    #[serde(skip_serializing)]
    id: i64,
    name: String,
    version: Option<String>,
}

#[derive(Serialize, FromRow)]
pub struct Variation {
    #[serde(skip_serializing)]
    id: i64,
    #[serde(skip_serializing)]
    package_id: i64,
    #[serde(skip_serializing)]
    distro_id: i64,
    name: Option<String>,
    version: Option<String>,
    package_url: String,
    download_url: String,
}

#[derive(Serialize, FromRow)]
pub struct FullVariation {
    #[serde(skip_serializing)]
    id: i64,
    #[serde(skip_serializing)]
    package_id: i64,
    #[serde(skip_serializing)]
    distro_id: i64,
    name: Option<String>,
    version: Option<String>,
    package_url: Option<String>,
    download_url: Option<String>,
    distro_name: String,
    distro_version: Option<String>,
}

// #[derive(Serialize, FromRow)]
// pub struct FullPackage {
//     #[serde(skip_serializing)]
//     id: i64,
//     package: Package,
//     name: String,
//     latest_version: String,
//     description: String,
//     keywords: Vec<String>,
//     homepage: String,
//     developer: Vec<String>,
// }
