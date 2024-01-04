use serde::Serialize;
use sqlx::{FromRow, Pool, Postgres};

pub struct AppState {
    pub(crate) db: Pool<Postgres>,
}

#[derive(Serialize, FromRow)]
pub struct Package {
    id: i64,
    human_name: String,
    name: String,
    latest_version: String,
    description: String,
    keywords: Vec<String>,
    homepage: String,
    developer: Vec<String>,
}

#[derive(Serialize, FromRow)]
pub struct Distro {
    id: i64,
    name: String,
    version: String,
}

#[derive(Serialize, FromRow)]
pub struct Variation {
    id: i64,
    package_id: i64,
    distro_id: i64,
    name: String,
    version: String,
    package_url: String,
    download_url: String,
}