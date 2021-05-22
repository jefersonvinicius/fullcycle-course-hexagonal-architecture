## Arquitetura Hexagonal
Permite a criação de um software isolando as complexidades do négocio das complexidades técnicas. Resultando em uma maior facilidade ao testar, alterar e aumenter este software.

O "core" da aplicação fica reponsável pela regras de negócio, enquanto as depêndencias (banco de dados, libs, protocolos de comunição) se comunicam com o "core" através de adaptadores.