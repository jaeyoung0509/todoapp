## GO & svelte todo apps

### BE
#### ENT 
- orm 
```shell
go install entgo.io/ent/cmd/ent@latest
```
- create schema 
```shell
go run -mod=mod entgo.io/ent/cmd/ent new Todo
```
### fx
- origin: `effects` 함수형 프로그래밍에서 사이드 이펙트를 관리하는 도구를 의미하는 용어
- Fx is di runtime di system for Go
- wih fx you can
```
- reduce boilerplate in setting up your application
- eliminate global state in your application
- add new components and have them instantly accessible across the application
- build general purpose sharable modules that just work
```
#### brief
- 생성
- 제공자 등록: fx.Provide
- 호출자 등록: fx.Invoke
- 생명주기 관리:fx.Lifecycle
- 에러 핸들링: fx.Options
#### fx.Invoke
- 함수를 인자로 받아 함수를 애플리케이션 시작 시점에 호출 
- 함수의 매개변수는 `fx.Provide`를 사용해 등록된 생성 함수를 통해 만들어진 값
- 이를 통해, 함수의 매개변수로 필요한 의존성들이 자동으로 주입
```go
// example 
fx.Invoke(handler.RegisterRoutes)로 작성하려면 
handler.RegisterRoutes 함수가 fx.Provide로 등록된 생성 함수가 만들어내는 타입들을 매개변수로 가져야 함
```
- 익명함수 사용 vs 직접 전달 
  -  둘의 차이는 함수의 시그니처가 `fx.Provide`로 등록된 생성 함수가 만드는 타입과 일치해야 됨, 반면 익명함수를 사용하면 중간에 변환역할을 할 수 있음
-> 즉 `fx.Invoke`가 요구하는 시그니처와 실제 함수의 시그니처가 다를 경우에도 사용할 수 있음

#### in giant software
- `fx.Options`를 사용하여 관련된 제공자와 호출자를 그룹화하고 모듈화 할 수 있음.
```go
func NewDatabaseModule() fx.Option {
    return fx.Options(
        fx.Provide(NewDatabaseConnection),
        fx.Provide(NewUserRepository),
        fx.Provide(NewProductRepository),
    )
}

func NewServerModule() fx.Option {
    return fx.Options(
        fx.Provide(NewServer),
        fx.Invoke(RegisterHTTPHandlers),
    )
}

func main() {
    fx.New(
        NewDatabaseModule(),
        NewServerModule(),
    )
}

```

### graceful shutdown
#### hooks
- hooks란 어떤 작업이나 이벤트에 대한 후속 처리를 위한 기능에 많이 쓰임
- 이는 특정 시점에 필요한 로직을 실행시키기위해 ㅁ낳이 사용됨
- Fx에서 hook은 애플리케이션 시작, 종료에 필요한 동작을 정의하는데 사용됨

#### gorouitine && channels
- 채널은 데이터를 전달하거나 실행 흐름을 제어하는 방법으로
데이터를 보내거나 받는 동작이 블로킹되는 특성이 있음
- 채널을 여러 고루틴간에 안전하게 주고받을 수 있음
- 고루틴이 채널에서 무한정 대기하는것은 오버헤드라고 생각될 수 있지만,
고루틴이 채널에서 대기하고 있는동안, go runtime은 다른 고루틴을 실행시키게 됨으로 cpu를 효율적으로 사용할 수 있음
-> 즉 하나으 고루틴이 대기상태에 들어가도, 다른 고루틴들이 계속 실행되어 시스템의 전체적인 성능에 영향을 주지 않음
- 채널에서 데이터를 받는 동작 (`data <-c`) + 버퍼를 1로 설정해, 채널에 데이터(signal)을 받을때까지 대기하게 됨 -> 이 때 고루틴은 블로킹 상태가 됨

#### Onstart 
- 고루틴을 사용하는 이유: 웹 서버외에도 다른 여러 작업들이 동시에 진행되어야 되기 때문
백그라운드에서 로그를 처리하거나, db를 확인하거나, 서드 파티와 통신
- 이러한 작업들은 웹 서버의 작업과 독립적으로 동시에 실행되어야 함
#### Onstop
- 메모리 누수 관점
- 고루틴은 독립적으로 실행되는 흐름으로, 한 고루틴에서 발생하는 이슈나 차단이 다른 고루틴에 영향을 주지 않음
이는 고루틴이 메인 고루틴과 독립적으로 실행되기 때문
- 이는 고루틴이 병렬로 실행할 수 있도록 여러 작업을 동시에 처리할 수 있게 해줌
- 한 작업이 차단되더라고 다른 작업은 계속 할 수 있음
- 하지만 이러한 특성 때문에 만약작업이 완료된 후에 실행되어야 하는 코드를 고루틴으로 만들면
그 코드는 다른 코루틴에 의해 차단될 수 있음.
그 코드가 실행되는 동안 다른 고루틴은 계속 실행되므로, 그 코드는 다른 고루틴의 상태에 따라 영향을 받을 수 있음
- 타이밍에 의존하지 않고, 어떠한 상황에서도 안전하게 실행


## thought archive
- adapter(repository)에서 도메인 영역으로 값을 전달할 때 domain(entity)에 맞춰서 전달해도 될까?
```
예시)
repository(adapter)영역에서 ent.Todo -> domain.Todo로 domain에 맞춰서 데이터를 전달해도 될 지 ?
-> adapter 영역이  domain의 구체적인 entity를 알아도 될 지?
```
- **col**: 헥사고날 아키텍처의 핵심 원칙 중 하나는 내부 레이어가 외부 레이어에 의존하면 안 됨. 여기서 내부 외부는 비지니스 로직 중심에서 떨어진 정도
  - 도메인 레이러는 아키텍처의 중심, infra 레이어는 외부
  - 도메인 레이어는 인프라 영역에 의존하면 안 됨
  - **인프라 레이어가 도메인 레이어를 허용하는것은 가능**



## Docker image optimization
### GCO_ENABLED=0
### Multi stage 
```yaml
# Start a new stage to reduce image size 
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the binary from the previous stage
COPY --from=builder /app/main /main

# Expose the port for the app to run on
EXPOSE 3000

# Run the binary 
CMD ["/main"]
```