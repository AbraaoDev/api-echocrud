## 🤝 EchoCRUD

Bem-vindo ao projeto da EchoCRUD! Este repositório contém o código-fonte de uma API RESTful, projetada para gerenciar Estabelecimentos e Lojas.

> [!IMPORTANT]
> 🔄 Testes Unitários
>
> 🔄 SOLID
>
> 🔄 Docker (Frontend + Backend)


## 🎯 Endpoints

### Establishments

|         Endpoint        |                    Router                    |                 Description                 |
|:-----------------------:|:--------------------------------------------:|:-------------------------------------------:|
| **[`get`](#get)**       | `/establishments`                            | Buscar todos os estabelecimentos            |
| **[`post`](#post)**     | `/establishment`                             | Criar um novo estabelecimento               |
| **[`get`](#get)**       | `/establishment/{establishmentId}`           | Buscar estabelecimento por ID               |
| **[`put`](#put)**       | `/establishment/{establishmentId}`           | Atualizar estabelecimento existente         |
| **[`delete`](#delete)** | `/establishment/{establishmentId}`           | Deletar estabelecimento por ID              |

### Stores

|         Endpoint        |                    Router                    |                 Description                 |
|:-----------------------:|:--------------------------------------------:|:-------------------------------------------:|
| **[`get`](#get)**       | `/establishments/{establishmentId}/stores`   | Buscar todas as lojas de um estabelecimento |
| **[`post`](#post)**     | `/establishments/{establishmentId}/stores`   | Criar nova loja em um estabelecimento       |
| **[`get`](#get)**       | `/stores/{storeId}`                          | Buscar loja por ID                          |
| **[`put`](#put)**       | `/stores/{storeId}`                          | Atualizar loja existente                    |
| **[`delete`](#delete)** | `/stores/{storeId}`                          | Deletar loja por ID                         |


## 🧪 Technologies[Front-End]

This project was developed using the following technologies:

- [VueJs](https://vitejs.dev/)
- [Shadcn-Vue](https://www.shadcn-vue.com/)
- [TypeScript](https://www.typescriptlang.org/)
- [Tailwind](https://tailwindcss.com/)
- [Zod](https://github.com/colinhacks/zod)


## 🧪 Technologies[Back-End]

- [Echo(GO)](https://echo.labstack.com/)
- [gORM](https://gorm.io/)
- [PostgreSQL](https://gorm.io/)

## 📝 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for details.

---

<p align="center">Made with ❤️ by Abraão DEV</p>
