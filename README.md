# GOLANG

## GOLANG의 특징
- 구글이 2009년 개발한 프로그래밍 언어로 빠른 속도와 안정성, 간단한 문법이 특징이다.
- C언어의 do, do while, while, for등의 반복문들을 for문 하나로 간략화 하였다.
- 컴파일 속도와 수행속도가 빠른 장점을 가진다.
- GOLANG은 컴파일 시 자료형이 결정되는 정적 타입이다.

### Benchmark
#### GOLANG
| Batch Size | Max Latency (s) | Avg Latency (s) | Throughput (TPS) |
|:----------|:-------------:|:------:| :-----:|
| 1 |  1.65 | 1.06 | 412.2 |
| 10 | 4.11 | 2.69 | 144.1 |
| 20 | 5.69 | 3.48 | 106.8 |
| 30 | 10.21 | 5.44 | 60.1 |
| 50 | 16.03 | 7.97 | 40.7 |

#### Javascript
| Batch Size | Max Latency (s) | Avg Latency (s) | Throughput (TPS) |
|:----------|:-------------:|:------:| :-----:|
| 1 |  1.17 | 0.41 | 363.2 |
| 10 | 4.03 | 1.93 | 53.4 |
| 20 | 5.80 | 2.93 | 35.3 |
| 30 | 9.80 | 4.80 | 21.0 |
| 40 | 11.73 | 5.22 | 17.4 |
| 50 | 15.13 | 6.42 | 13.4 |
> block size : 8k byte

## Heading
---------------------------
* 다른 패키지에서 사용가능한 함수는 대문자 사용, 소문자로 시작하면 build-in/private 함수
* := 는 할당연산자로 var를 명시하지 않고 변수 선언과 초기화를 동시에 수행
* golang은 couchbase, mongodb와 잘맞음
* golang의 특징으로 class 기반이 아니기 때문에 method의 모음은 interface이고 data의 모음은 structure 라고 할수있음
* channel은 프로세스간 통신을 위함
* import ."fmt" // .이붙을 경우 Println("")으로 사용가능, package가 여러 개고 각 package끼리 의존적일 경우 .사용은 충돌을 발생시킬 수 있음
* 그냥 대괄호만 지정하면 slice / [...] 이나 숫자가 있다면 배열
* 배열은 하나하나씩 값을 세팅하지만 slice는 한번에 값을 세팅할 수 있음
* nil은 null과 유사함
* 함수내에서 내장함수선언은 익명함수만 가능함 이름을 가지면 에러발생
* GOLANG 에도 여러 Web Framework가 존재