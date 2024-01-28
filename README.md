# packsearch
A new way to search Linux packages across distributions

This project is still at the design phase.

## Why?
Being a Fedora user, I sometimes want to compile software on my computer. However, sometimes, the compile instructions only give the dependencies for an Ubuntu system, and I have to go hunting for the equivalent package names for my distro. This project aims to solve this.

## How?
I intend to build an API and a website that make it easy to find and search packages from a keyword or list of packages.

At some point in time, I want this to be more than just an API and a website, I want to also have a CLI utility that would allow to install a package easily regardless of the distro.

For example, you could run `cmd install --arch ${arch-package-name}` on an Ubuntu distro, and it would fetch the API, and with the response, automatically download the right package from the Ubuntu repository.
With this, I want to make it easier for people to compile software so that on installation instructions, the projects can simply use this tool for the dependencies part. Lets say the project gives Ubuntu compile instructions, they could easily switch their command code block to use this program with the --ubuntu flag specified so that the users do not have to go hunting for the equivalent packages on their distro.

---

## API

Lets start by defining the terms that are going to be used here:
- Package: The top-level and non-changing information of a package.
- Variation: The variation of a package for a certain distribution. This contains the data that varies from distro to distro.

Note: the routes, terms and params could change if better ones are found, this is mostly to have a guideline of what to create and prevent too much deviation

### /search/
param: one package name to look for
Would return a list of packages with their "formal" and formatted names. Under this object, there would be the basic information of the package like the developper(s), description, links, etc. There would ideally also be an array/list that contains the equivalent name for the various distros,
```json
{
  "id" : 1,
  "human_name" : "Java Open JDK 17",
  "name" : "java_open_jdk_17",
  "description" : "OpenJDK 17 Runtime Environment",
  "developer" : "OpenJDK",
  "homepage" : "https://openjdk.org/projects/jdk/17/",
  "variations" : [
    {
      "ubuntu" : {
        "name" : "openjdk-17-jdk",
        "version" : "17.0.9"
      },
      "fedora" : {
        "name" : "java-17-openjdk",
        "version" : "17.0.9"
      }
    }
  ]
}
```

### /search-list/
param: multiple package names to look for, divided by a separator
This would be the same as /search/, but with a list of multiple different packages instead of just one.
```json
{
  [
    {
      "id" : 1,
      "human_name" : "Java Open JDK 17",
      "name" : "java_open_jdk_17",
      "description" : "OpenJDK 17 Runtime Environment",
      "developer" : "OpenJDK",
      "homepage" : "https://openjdk.org/projects/jdk/17/",
      "variations" : [
        {
          "ubuntu" : {
            "name" : "openjdk-17-jdk",
            "version" : "17.0.9"
          },
          "fedora" : {
            "name" : "java-17-openjdk",
            "version" : "17.0.9"
          }
        }
      ]
    },
    {
      "id" : 2,
      "human_name" : "Java Open JDK 17",
      "name" : "java_open_jdk_17",
      "description" : "OpenJDK 17 Runtime Environment",
      "developer" : "OpenJDK",
      "homepage" : "https://openjdk.org/projects/jdk/17/",
      "variations" : [
        {
          "ubuntu_2310" : {
            "name" : "openjdk-17-jdk",
            "version" : "17.0.9"
          },
          "fedora_39" : {
            "name" : "java-17-openjdk",
            "version" : "17.0.9"
          }
        }
      ]
    }
  ]
}
```

### /package/:package_name/add-variation/:variation/
:package_name -> the name (or ID, TDB later) of the package to add a variation to.
:variation -> The variation of the package to add to the package.
This endpoint would add a variation to a top-level package.

## Database
I would ideally want this to be a relational (SQL) database because in this case, relations could be very useful.

### Schema
#### Distro
There would need to be a distro table, to keep track of distro versions and to make sure there are no conflicts (if for example ubuntu 22.04 has a different package as 24.04). Rolling release distros and other means of distribution that do not have a version tied to package versions such as Flatpak, Snap or Homebrew are also supported, simply with a null in the 'version' column.

| Name    | Type | Nullable? | Description                                        |
|---------|------|-----------|----------------------------------------------------|
| id      | int8 | No        |                                                    |
| name    | text | No        | Name of the distribution method.                   |
| version | text | Yes       | Version of the distribution method, if applicable. |

##### Example data
| id | name    | version |
| -- | ------- | ------- |
| 1  | Ubuntu  | 23.10   |
| 2  | Ubuntu  | 24.04   |
| 3  | Fedora  | 39      |
| 4  | Arch    | NULL    |
| 5  | Flathub | NULL    |

#### Variation
The variation table contains all the variations across the distribution methods.
| Name         | Type | Nullable? | Description                                               |
|--------------|------|-----------|-----------------------------------------------------------|
| id           | int8 | No        |                                                           |
| package_id   | int8 | No        | ID of the package.                                        |
| distro_id    | int8 | No        | ID of the distribution method.                            |
| name         | text | Yes       | Name of the package in the repositories of the distro.    |
| version      | text | Yes       | Version of the package.                                   |
| package_url  | text | Yes       | URL of the package page for the repository of the distro. |
| download_url | text | Yes       | Direct download URL of the package for the distro.        |

Example data:
| package_id | distro_id | name            | version |
| ---------- | --------- | --------------- | ------- |
| 1          | 1         | openjdk-17-jdk  | 17.0.9  |
| 1          | 2         | java-17-openjdk | 17.0.9  |

#### Package
I am currently thinking of something like the following as the main table without the distro-specific data. (note: the keywords column could be a separate table)

| Name           | Type   | Nullable? | Description                                                    |
|----------------|--------|-----------|----------------------------------------------------------------|
| id             | int8   | No        |                                                                |
| human_name     | text   | No        | Human-readable name of the package.                            |
| name           | text   | Yes       |                                                                |
| latest_version | text   | Yes       | Latest available version of the package.                       |
| description    | text   | Yes       | Developer's description of the package.                        |
| keywords       | text[] | Yes       | Keywords associated with the package. Used for search queries. |
| homepage       | text   | Yes       | URL of the homepage of the package.                            |
| developer      | text[] | No        | Developers or contributors of the project.                     |

Example data:
| id | human_name     | name             | description                    | keywords                   | homepage                                | latest_version | developer |
|----|----------------|------------------|--------------------------------|----------------------------|-----------------------------------------|----------------|-----------|
| 1  | Java Open JDK  | java_open_jdk_17 | OpenJDK 17 Runtime Environment | java,openjdk,jdk,jdk17,... | https://openjdk.org/projects/jdk/17/    | 17.0.9         | OpenJDK   |
| 2  | Firefox        | firefox          | Mozilla Firefox Web browser    | NULL                       | https://www.mozilla.org/fr/firefox/new/ | 122            | Mozilla   |

