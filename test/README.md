# Tipo de teste

## Smoke testing
O teste de fumaça é um teste de carga regular, configurado para carga mínima. Você deseja executar um teste de fumaça como uma verificação de sanidade toda vez que escrever um novo script ou modificar um script existente.

Você deseja executar um teste de fumaça para:
    1 - Verifique se seu script de teste não contém erros.
    2 - Verifique se o seu sistema não gera nenhum erro quando está sob carga mínima.

## Load testing
O Load Testing se preocupa principalmente com a avaliação do desempenho atual do seu sistema em termos de usuários ou solicitações simultâneas por segundo.

Quando você quiser entender se seu sistema está atendendo às metas de desempenho, esse é o tipo de teste que você executará.

O que é teste de carga
Teste de carga é um tipo de teste de desempenho usado para determinar o comportamento de um sistema em condições normais e de pico.

O Load Testing é usado para garantir que o aplicativo tenha um desempenho satisfatório quando muitos usuários o acessam ao mesmo tempo.

Você deve executar um teste de carga para:
    1 - Avalie o desempenho atual do seu sistema sob carga típica e de pico.
    2 - Certifique-se de continuar atendendo aos padrões de desempenho ao fazer alterações em seu sistema (código e infraestrutura).

## Stress testing
Teste de estresse é um tipo de teste de carga usado para determinar os limites do sistema. O objetivo deste teste é verificar a estabilidade e confiabilidade do sistema sob condições extremas .

Para executar um teste de estresse adequado, você precisa de uma ferramenta para empurrar o sistema acima de suas operações normais, até seus limites e além do ponto de ruptura .

Normalmente, você deseja fazer um teste de estresse em uma API ou site para determinar:
    1 - Como seu sistema se comportará sob condições extremas.
    2 - Qual é a capacidade máxima do seu sistema em termos de usuários ou taxa de transferência.
    3 - O ponto de ruptura do seu sistema e seu modo de falha.
    4 - Se o seu sistema se recuperar sem intervenção manual após o término do teste de estresse.

Ao testar o estresse, você configurará o teste para incluir mais usuários simultâneos ou gerar uma taxa de transferência maior do que:
    1 - Seu aplicativo normalmente vê.
    2 - Você acha que vai ser capaz de lidar.

## Soak testing
O teste de imersão está preocupado com a confiabilidade por um longo período de tempo. Os problemas de confiabilidade geralmente estão relacionados a bugs, vazamentos de memória, cotas de armazenamento insuficientes, configuração incorreta ou falhas de infraestrutura. Os problemas de desempenho geralmente estão relacionados ao ajuste incorreto do banco de dados, vazamentos de memória, vazamentos de recursos ou uma grande quantidade de dados.

Com um teste de imersão, você pode simular dias de tráfego em apenas algumas horas.

Você normalmente executa este teste para:
    1 - Verifique se o seu sistema não sofre de bugs ou vazamentos de memória, que resultam em travamento ou reinicialização após várias horas de operação.
    2 - Verifique se as reinicializações esperadas do aplicativo não perdem solicitações.
    3 - Encontre bugs relacionados a condições de corrida que aparecem esporadicamente.
    4 - Certifique-se de que seu banco de dados não esgote o espaço de armazenamento alocado e pare.
    5 - Certifique-se de que seus logs não esgotem o armazenamento em disco alocado.
    6 - Certifique-se de que os serviços externos dos quais você depende não parem de funcionar após a execução de uma determinada quantidade de solicitações.

fonte: https://k6.io/docs/test-types/introduction/