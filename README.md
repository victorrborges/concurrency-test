# Avaliação 3 - Programação Concorrente

Repositório contendo as soluções para as questões da terceira avaliação da disciplina Programação Concorrente (UFCG - 2020.0)

Aluno: Victor Eduardo Borges de Araújo

Email: victor.araujo@ccc.ufcg.edu.br

Matrícula: 115210597

# Primeira Questão

Q1 - Suponha que exista um conjunto de n threads passageiras e uma thread carro. As threads passageiras, repetidamente, espera para pegar uma carona no carro, o qual pode ser ocupado por C passageiros (C < n). O carro só pode partir quando estiver completo. Escreva o código das abstrações passageiro e carro. Considere os detalhes adicionais abaixo:


  - As threads passageiras devem chamar os métodos/funções embarcar e desembarcar;
  - A thread carro pode executar os métodos/funções carregar, correr e descarregar;
  - Passageiros não podem embarcar até que o carro tenha chamado  executado carregar;
  - O carro não poder correr até que todos os C passageiros tenham embarcado;
  - Os passageiros não podem desembarcar até que o carro tenha executado descarregar.

Como executar a solução da primeira questão:

```sh
sh executar_questao_1.sh 
```

# Segunda Questão

Q2 - Na beira do açude de bodocongó existe um barco de um remo só. Tanto os alunos de computação da UFCG quanto os alunos de psicologia da UEPB usam o barco para cruzar o açude. O barco só pode ser usado por exatamente quatro alunos; nem mais, nem menos. Para garantir a segurança dos discentes, é proibido que um aluno da UFCG seja colocado no barco com três alunos da UEPB. Também é proibido colocar um aluno da UEPB com três alunos da UFCG. Os alunos são threads. Para embarcar, cada thread chama a função/método embarcar. Você precisa garantir que toda as quatro threads de um carregamento de alunos chamam a função embarcar antes de qualquer outra threads, de um próximo carregamento. Após todas as quatro threads de um carregamento tenham chamado a função embarcar, uma única thread dessas quatro (não importa qual delas seja) deve executar a função rema, indicando que essa thread será responsável por assumir o papel de remadora. Assuma que depois que o barco chega ao destino, magicamente ele reaparece na origem (ou seja, assuma que estamos interesados no tráfego do barco em somente um sentido). Implemente as threads AlunoUFCG e AlunoUEPB bem como qualquer código utilitário para compor sua solução.


Como executar a solução da segunda questão:

```sh
sh executar_questao_2.sh 
```


# Terceira Questão

Q3 - Os alunos do período 2003.1 adoravam beber no bar de Auri. Esse era um bar bastante movimentado. Então, a única mesa do bar era compartilhada por desconhecidos. As threads Aluno podem chamar as funções/métodos bebe e depois disso, sai. Após chamar a função bebe e antes de chamar a função sai,  o Aluno é considerado remediado. Existia uma única regra no bar de Auri: ninguém pode ser deixando bebendo sozinho na mesa. Um aluno bebe sozinho se todo mundo que estava compartilhando a mesa com ele, ou seja, todo os demais que chamaram a função bebe, chamam a função sai antes do bebedor terminar de executar sua função bebe. Escreva o código das threads Aluno e qualquer outro código utilitário que garanta a restrição de não deixar nínguem bebendo sozinho (e, obviamente, apresente progresso, ou seja, que permita que os Alunos bebam).

Como executar a solução da terceira questão:

```sh
sh executar_questao_3.sh 
```
