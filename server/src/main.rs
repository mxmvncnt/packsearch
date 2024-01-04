use actix_web::{App, get, HttpResponse, HttpServer, Responder, web::Data};
use dotenv::dotenv;
use serde::Serialize;
use sqlx::{FromRow, Pool, Postgres, postgres::PgPoolOptions};

pub struct AppState {
    db: Pool<Postgres>,
}

#[derive(Serialize, FromRow)]
struct Package {
    id: i64,
    human_name: String,
    name: String,
    latest_version: String,
    description: String,
    keywords: Vec<String>,
    homepage: String,
    developer: Vec<String>
}

#[get("/")]
async fn hello(state: Data<AppState>) -> impl Responder {
    let result = sqlx::query_as::<_, Package>(
        "
            SELECT
                id,
                human_name,
                name,
                latest_version,
                description,
                keywords,
                homepage,
                developer
            FROM package;
            ")
        .fetch_all(&state.db)
        .await;

    match result {
        Ok(packages) => HttpResponse::Ok().json(packages),
        Err(error) => {
            println!("{}", error.to_string());
            HttpResponse::NotFound().json("No packages found.")
        }
    }
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    dotenv().ok();

    // base connection code from https://github.com/bocksdin/rust-sqlx/tree/main
    let database_url = std::env::var("DB_URL").expect("DB_URL must be set");
    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&database_url)
        .await
        .expect("Error building a connection pool");

    HttpServer::new(move ||
        App::new()
            .app_data(Data::new(AppState { db: pool.clone() }))
            .service(hello)
    )
        .bind(("127.0.0.1", 8080))?
        .run()
        .await
}
