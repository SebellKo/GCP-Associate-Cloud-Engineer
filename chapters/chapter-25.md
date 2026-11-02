# Chapter 25. 서비스 계정·Impersonation·Workload Federation

[← 이전 장](chapter-24.md) · [전체 목차](../README.md) · [다음 장 →](chapter-26.md)

공식 가이드 연결: **4.2** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 사람과 앱 identity를 분리한다
서비스 계정은 앱·자동화의 identity이며 그 자체도 IAM 정책을 가질 수 있는 resource다. 앱이 버킷을 읽는 권한과 개발자가 그 서비스 계정을 사용하는 권한은 별개다. Service Account User는 actAs 용도, Token Creator는 impersonation용 단기 토큰 생성 등의 용도다. 두 역할이 모든 상황에서 서로 대체된다고 생각하지 않는다.

Google Cloud 안에서는 실행 자원에 전용 서비스 계정을 연결하고 단기 자격 증명을 사용한다. 외부 CI/CD는 Workload Identity Federation으로 외부 OIDC 같은 신뢰를 연결해 장기 JSON 키를 줄인다. issuer·audience·attribute condition을 특정 repo·branch·environment 요구와 맞춰 제한한다.

GKE Workload Identity Federation은 Pod별 identity로 Google API 접근을 구성한다. 이미지 pull identity와 다르다. Workforce Federation은 직원 등 사람의 외부 IdP 연결, Workload Federation은 앱·작업 identity 연결이다. Google-managed service agent는 제품 작업을 수행하므로 이름만 보고 권한을 제거하지 않는다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. actAs와 token 생성은 같은 permission인가?

2. Workforce/Workload의 대상은?

3. WIF 설정에서 repository 조건이 필요한 이유는?

## 3. 실습과 관찰

통합 실습의 runtime service account를 조회한다. 정책에서 resource 접근 binding과 계정 자체의 impersonation binding을 구분한다. 외부 CI의 federation은 issuer·audience·repo 조건 설계까지만 기본으로 한다. 테스트 계정은 연결 자원을 지운 뒤 삭제하며 JSON private key는 만들지 않는다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

외부 GitHub Actions에서 GCP에 배포하되 장기 키는 저장하지 않는다.

A. Owner JSON 키를 공개 저장소에 둔다
B. 모든 bucket을 공개한다
C. OIDC 기반 Workload Identity Federation을 구성한다
D. 개발자 비밀번호를 공유한다

### 문제 2 — 단일 선택

개발자가 특정 서비스 계정을 impersonate해 단기 토큰으로 작업해야 한다.

A. 해당 계정의 필요한 Token Creator 권한과 자원 권한을 구성한다
B. Storage 등급만 바꾼다
C. DNS A 레코드를 만든다
D. 개인 SSH 키만 등록한다

### 문제 3 — 단일 선택

외부 기업 IdP의 직원들이 Google Cloud 콘솔에 접근해야 한다.

A. Cloud NAT로 인증을 대신한다
B. Spot VM을 생성한다
C. Pub/Sub topic만 만든다
D. Workforce Identity Federation을 검토한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. 사람/앱·자동화다.

3. 의도한 외부 workload로 신뢰를 제한한다.

### 시험형 예제

**1. C** — 외부 workload를 단기 federation으로 인증한다. 나머지는 비밀 노출 또는 권한 과다다.

**2. A** — 토큰 생성 권한과 계정이 수행할 자원 권한이 필요하다. Storage·SSH·DNS는 impersonation 구성이 아니다.

**3. D** — 직원 identity federation은 Workforce 영역이다. 메시징·NAT·VM은 사람의 SSO를 대체하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/iam/docs/workload-identity-federation)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-24.md) · [다음 장 →](chapter-26.md)
