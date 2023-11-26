## AFEB API

Esta é a API oficial do website da AFEB. Aqui você encontra todos os endpoints
do sistema para colaborar no desenvolvimento deste projeto, ou então conectar-se
a ele para a criação de um projeto seu.

Link da API: <a href="https://afeb-api.onrender.com/">https://afeb-api.onrender.com/</a>

# Endpoints

* [/api/jogadores](#apijogadores)

## /api/jogadores

<details>
  <summary> <code>GET</code> <code>/api/jogadores/ranking</code> </summary>

  ### Descrição

  Retorna um ranking com o Top 10 jogadores mais bem rankeados da AFEB.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "ranking": []object,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Ranking encontrado com sucesso!",
    "ranking": [
      {
        "codJog": 1,
        "nome": "Lucas Guedes",
        "apelido": "Guedes",
        "tituloAFEB": "MNB",
        "info": "Fundador da AFEB.",
        "eloRapid": 3200,
        "eloBlitz": 3400,
        "dataNascimento": "2005-09-02",
        "trofeus": null
      },
      {
        "codJog": 2,
        "nome": "Bye",
        "apelido": null,
        "tituloAFEB": "GMB",
        "info": "O bye",
        "eloRapid": 9999,
        "eloBlitz": 9999,
        "dataNascimento": "0000-01-01",
        "trofeus": null
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/jogadores/${codJog}</code> </summary>

  ### Descrição

  Retorna os dados de um jogador, juntamente de suas premiações na AFEB.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codJog    |  required | string      | Código do jogador a ser buscado.                                 |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "jogador": object,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "jogador": {
      "codJog": 2,
      "nome": "Bye",
      "apelido": null,
      "tituloAFEB": "GMB",
      "info": "O bye",
      "eloRapid": 9999,
      "eloBlitz": 9999,
      "dataNascimento": "0000-01-01",
      "trofeus": [
        {
          "codTrof": 0,
          "codJog": 0,
          "codTorn": 0,
          "torneio": "1º Campeonato Mundial da FIDE",
          "posicao": 1
        }
      ]
    },
    "message": "Jogador encontrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/jogadores</code> </summary>

  ### Descrição

  Registra um novo jogador no sistema

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "nome": string,
    "apelido": string || null,
    "tituloAFEB": string || null,
    "info": string || null,
    "eloRapid": Number || null,
    "eloBlitz": Number || null,
    "dataNascimento": string
  }
  ```

  Exemplo:
  ```js
  {
    "nome": "Bye",
    "apelido": "Ciao",
    "tituloAFEB": "GMB",
    "info": "O bye.",
    "eloRapid": 9999,
    "eloBlitz": 9999,
    "dataNascimento": "1287-01-01"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador registrado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/jogadores</code> </summary>

  ### Descrição

  Atualiza os dados de um jogador no sistema

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codJog": Number,
    "nome": string,
    "apelido": string || null,
    "tituloAFEB": string || null,
    "info": string || null,
    "eloRapid": Number || null,
    "eloBlitz": Number || null,
    "dataNascimento": string
  }
  ```

  Exemplo:
  ```js
  {
    "codJog": 0,
    "nome": "Bye",
    "apelido": "Ciao",
    "tituloAFEB": "GMB",
    "info": "O bye.",
    "eloRapid": 9999,
    "eloBlitz": 9999,
    "dataNascimento": "1287-01-01"
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador atualizado com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/jogadores/${codJog}</code> </summary>

  ### Descrição

  Exclui um jogador permanentemente do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codJog    |  required | string      | Código do jogador a ser excluído.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Jogador excluído com sucesso!"
  }
  ```
</details>

## /api/noticias

<details>
  <summary> <code>GET</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Retorna todas as notícias da AFEB.

  > Em desenvolvimento
</details>

<details>
  <summary> <code>GET</code> <code>/api/noticias/feed</code> </summary>

  ### Descrição

  Retorna as 6 últimas notícias da AFEB.

  ### Parâmetros

  > Nenhum

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "noticias": []object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícias encontradas com sucesso!",
    "noticias": [
      {
        "codNotc": 1,
        "codAutor": "OGrandePoderosoB",
        "titulo": "AFEB continua existindo",
        "noticia": "Atualização para o dia de hoje: A AFEB ainda existe.",
        "dataPublicacao": "2023-11-24 11:23:50"
      }
    ]
  }
  ```
</details>

<details>
  <summary> <code>GET</code> <code>/api/noticias/${codNotc}</code> </summary>

  ### Descrição

  Retorna uma notícia específica do jornal da AFEB a partir de seu código.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codNotc   |  required | string      | Código de notícia a ser retornada.                               |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "noticia": object || null,
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia encontrada com sucesso!",
    "noticia": {
      "codNotc": 7,
      "codAutor": "LucasMoraesGuede",
      "titulo": "Murilo Holtz se abstém do Torneio de Chess 960",
      "noticia": "A partir de hoje, quinta-feira (16), Murilo Holtz não irá mais participar do Torneio de Chess 960, pois irá fazer uma viagem aos Estados Unidos, e não será capaz de comparecer ao torneio até sexta-feira da próxima semana (24). Assim, o torneio agora acontecerá apenas entre 8 participantes e com um forte competidor a menos.",
      "dataPublicacao": "2023-11-16 17:38:59"
    }
  }
  ```
</details>

<details>
  <summary> <code>POST</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Cria uma nova notícia.

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 201 | `CREATED` |
  | 400 | `BAD REQUEST` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codAutor": string,
    "titulo": string,
    "noticia": string
  }
  ```

  Exemplo:
  ```js
  {
    "codAutor": "OGrandePoderosoB",
    "titulo": "AFEB continua existindo",
    "noticia": "Atualização para o dia de hoje: A AFEB ainda existe."
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia criada com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>PUT</code> <code>/api/noticias</code> </summary>

  ### Descrição

  Atualiza os dados de uma notícia

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Request body
  ```js
  {
    "codNotc": Number,
    "codAutor": string,
    "titulo": string,
    "noticia": string
  }
  ```

  Exemplo:
  ```js
  {
    "codNotc": 2,
    "codAutor": "OGrandePoderosoB",
    "titulo": "AFEB continua existindo",
    "noticia": "Atualização para o dia de hoje: A AFEB ainda existe."
  }
  ```

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia atualizada com sucesso!"
  }
  ```
</details>

<details>
  <summary> <code>DELETE</code> <code>/api/noticias/${codNotc}</code> </summary>

  ### Descrição

  Exclui uma notícia do sistema.

  ### Parâmetros

  | param.    |  tipo     | data type   | desc.                                                            |
  |-----------|-----------|-------------|------------------------------------------------------------------|
  | codNotc   |  required | string      | Código da notícia a ser excluída.                                |

  ### Status codes

  | Status Code | Description |
  | :--- | :--- |
  | 200 | `OK` |
  | 400 | `BAD REQUEST` |
  | 404 | `NOT FOUND` |
  | 500 | `INTERNAL SERVER ERROR` |

  ### Response
  ```js
  {
    "message": string || null,
    "error": string || null
  }
  ```

  Exemplo:
  ```js
  {
    "message": "Notícia excluída com sucesso!"
  }
  ```
</details>