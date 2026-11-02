# Chapter 17. Cloud Run 배포·인증·리비전

[← 이전 장](chapter-16.md) · [전체 목차](../README.md) · [다음 장 →](chapter-18.md)

공식 가이드 연결: **2.1 / 3.1** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 요청 중심 실행 환경
Cloud Run service는 HTTP 요청을 처리하는 컨테이너를 관리한다. 컨테이너는 지정된 PORT로 수신하고 인스턴스 수명에 의존하는 로컬 상태를 영구 저장소로 취급하지 않는다. jobs는 요청 서버와 다른 작업 실행 모델이다. OS 커널 제어나 Kubernetes API가 필수라면 다른 실행 환경을 검토한다.

배포는 revision을 만든다. 새 revision에 처음부터 트래픽을 주지 않고 검증한 뒤 비율을 늘리거나 이전 revision으로 되돌릴 수 있다. 트래픽 롤백이 DB 스키마나 외부 상태까지 되돌리지는 않는다.

호출자 인증과 runtime service account를 구분한다. Cloud Run Invoker는 서비스를 부르는 권한, 런타임 계정의 DB 권한은 서비스가 외부 자원을 쓰는 권한이다. min instances는 지연·비용, max instances와 concurrency는 부하·DB 연결 수에 영향을 준다. max 설정을 모든 과금의 정확한 하드 캡으로 보지 않는다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. 호출자 권한과 런타임 계정 권한은 같은가?

2. revision 롤백으로 DB 마이그레이션도 자동 복원되는가?

3. min instances를 늘리면 어떤 trade-off가 있는가?

## 3. 실습과 관찰

[통합 실습](../labs/capstone.md)에서 비공개 Cloud Run을 배포하고 ID token으로 호출한다. 두 버전을 배포해 revision 목록을 기록하고 트래픽 변경 화면을 확인한다. 실습 종료 시 서비스뿐 아니라 빌드 이미지·artifact repository도 정리한다. 실제 클라우드 실행은 이 저장소 업로드 과정에서 수행하지 않았다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

새 버전에 5%만 보내며 안정성을 검증하려 한다.

A. 새 revision에 트래픽을 분할한다
B. 기존 DB를 삭제한다
C. 서비스를 무조건 공개한다
D. 로그를 끈다

### 문제 2 — 단일 선택

Cloud Run 호출이 403이다. 호출자는 토큰 없이 접근했다.

A. bucket class를 변경한다
B. 런타임 DB 메모리만 늘린다
C. 인증 설정·ID token·호출자의 Invoker 권한을 점검한다
D. CPU quota만 증액한다

### 문제 3 — 단일 선택

Cloud Run이 DB에 접근할 때 사용할 계정은?

A. 모든 호출자의 개인 Owner 키를 이미지에 넣는다
B. allUsers에 DB 관리자 권한을 준다
C. DNS 레코드로 계정을 대신한다
D. 전용 runtime service account에 필요한 DB 권한만 부여한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. 아니다.

3. 준비된 인스턴스와 지연 개선 대 비용 증가다.

### 시험형 예제

**1. A** — revision별 트래픽 분할이 점진 배포 기능이다. 공개화·DB 삭제·로그 중지는 검증을 대체하지 않는다.

**2. C** — 서비스 호출 인증과 권한 문제를 본다. DB 사양·CPU quota·저장 등급은 호출자 인증을 해결하지 않는다.

**3. D** — 런타임 identity를 최소 권한으로 구성한다. 개인키·공개 관리자·DNS는 안전한 실행 인증 방식이 아니다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/run/docs/overview/what-is-cloud-run)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-16.md) · [다음 장 →](chapter-18.md)
