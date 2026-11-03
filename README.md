Mais sobre a ferramenta: https://valterbarros.github.io/articles/blog-editor/

Esse projeto veio de outro repo, o historico completo do git pode ser encontrado em: https://github.com/valterbarros/valterbarros.github.io/compare/37e7e4ed552fb9f2bd63535412d1068b2878a67a...abbb7443ea39accf313b5388bf6baffcad6735a9

Tarefas

O projeto atual do blog é gerenciado usando jekyll e o gh pages, então a parte de templates, includes, variaveis de template está resolvido

Existe um workflow hoje que eu preciso fazer manualmente e ele seria:
- [x] criar um novo draft para um post, tudo que fica na pasta articles/drafts
- [x] criar um novo post, tudo que fica na pasta articles/_posts
- [x] converter um draft para um post ou vice versa
- [x] "publicar" um novo post
  - [x] fazer o commit e push
  - [x] Alterar nome do post para ter a data
  - [x] adicionar placeholder de metadata no topo do arquivo
  - [x] colocar tags de justificado em todos os paragrafos
  - [x] substituir tags de markdown como ** []() não funciona dentro do <p> é preciso substituir por tags html
- [x] listar todos os drafts
- [x] listar todos os posts
- [x] Abrir um post ou draft diretamente no vscode sem estar com ele aberto
- [x] Um liveview para ver o markdown do post enquanto digito poderia ser no navegador
  Etapas:
    - [x] Servir pagina html do editor/preview
    - [x] implementar websockets no server
    - [x] fazer parser do markdown 
- [x] Fazer envio de imagens ao colar imagens no editor 
  - [x] Servir as imagens do editor pasta estatica
  - [x] Fazer envio no frontend da imagem, e exibir a tag de imagem
  - [x] no preview exibir a imagem renderizada
  - [x] Criar endpoint ou websocket
  - [x] Salvar imagem na pasta articles/_posts/assets quando fier upload
  - [x] Exibir imagem no blog in live
  - [] Comprimir imagens
  - [x] Extrair cli para projeto a parte e integrar

Refatorações

- [x] Alguns metodos são iguais entre post e draft tentar generalizar eles 

Seria interessante:
- [x] Hoje eu preciso mandar o post para o google docs quando ele passa a ficar grande e complexo, se tivesse uma maneira barata de fazer isso seria interessante usando o cli claro
- [] Tornar o editor colaborativo onde outras pessoas possam editar junto como no docs
- [] Talves uma maneira de ir coletando ideias e dai ir colocando para depois ver o que realmente vale a pena e não

Esses pontos poderiam ser resolvidos usando uma ferramente de linha de comando

Fora do escopo

- Será que da pra substituir o jekyll? no caso fazer na mão mesmo
- Funcionalidades atuais do jekyll
  - Fazer include de uma pagina dentro de outra
  - Lidar com variaveis que ficam no topo da pagina
  - Arquivo de config com algumas configs
  - Montar paginas dinamicamentes, com for e ifs, um bom exemplo seria a index de posts que exibe todos os posts
  - Lidar com imports de assets estaticos imagens etc
  - Lidar com layouts

- Paginas do blog e pagina inicial
  - Pagina inicial do site em ptbr
  - Pagina inicial do site em ingles
  - Pagina de listagem de artigos
  - Pagina do artigo
