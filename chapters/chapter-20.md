# Chapter 20. Terraform·Helm·AI 보조 인프라 도구

[← 이전 장](chapter-19.md) · [전체 목차](../README.md) · [다음 장 →](chapter-21.md)

공식 가이드 연결: **2.4** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 원하는 상태를 코드로 표현한다
Terraform은 resource 선언과 state를 사용해 실제 자원의 변경 계획을 계산한다. `init`은 provider/backend 초기화, `validate`는 구성 검사, `plan`은 변경 검토, `apply`는 적용, `destroy`는 관리 자원 삭제다. state를 지운다고 실제 자원이 사라지지 않는다. 기존 자원은 import 등으로 관리 편입을 검토한다.

팀 state는 접근 제어·잠금·버전 관리를 갖춘 원격 backend를 사용한다. state와 plan에 비밀이 포함될 수 있으므로 공개 Git에 넣지 않는다. provider lock 파일은 재현 가능한 의존성 선택을 위해 관리한다. 일상적으로 -target에 의존하면 전체 의존 관계 변경을 놓칠 수 있다.

Helm은 Kubernetes 패키징, Config Connector는 Kubernetes 방식으로 Cloud 자원을 관리하는 도구다. Fabric FAST는 클라우드 기반 환경 구성을 돕는다. Gemini CLI·Cloud Assist·Antigravity·Application Design Center 같은 AI 도구는 계획과 코드 작성을 보조한다. 생성 결과도 diff·plan·권한·비용·정책을 검토한다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. plan과 apply는 같은가?

2. state 삭제가 자원 삭제인가?

3. AI가 만든 IaC를 검증할 항목은?

## 3. 실습과 관찰

[Terraform 실습](../labs/terraform.md)의 비공개 bucket 구성을 읽고 fmt/validate/plan을 수행한다. apply는 전용 학습 프로젝트에서만 한다. destroy 계획도 검토하고 완료 후 bucket 목록에서 삭제를 확인한다. state는 추적하지 않고 provider lock 파일만 의도적으로 관리한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

Terraform이 어떤 자원을 삭제할지 적용 전에 보고 싶다.

A. 무조건 apply -auto-approve를 실행한다
B. terraform plan으로 변경을 검토한다
C. 모든 자원을 수동 복사한다
D. state를 먼저 지운다

### 문제 2 — 단일 선택

공개 Git 저장소에 Terraform state를 넣으려 한다.

A. 확장자만 txt로 바꿔 공개한다
B. 민감 정보 위험을 고려해 보호된 state backend를 사용한다
C. bucket을 allUsers로 공개한다
D. 모든 비밀을 output으로 출력한다

### 문제 3 — 단일 선택

AI가 광범위한 IAM 권한과 DB 삭제를 포함한 plan을 제시했다.

A. 결제 export를 지운다
B. 로그만 끈다
C. AI 생성이므로 즉시 적용한다
D. 요구사항에 맞게 수정하고 plan·권한·삭제 영향을 검증한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 변경 검토와 실제 적용이다.

2. 아니다.

3. 계획·정책·권한·비용·삭제 영향이다.

### 시험형 예제

**1. B** — plan은 적용 전 검토다. state 삭제는 추적 손실이며 무검토 적용은 요구와 반대다.

**2. B** — state는 민감 데이터를 담을 수 있다. 이름 변경·output·공개화는 보호를 제공하지 않는다.

**3. D** — AI 제안도 검토 대상이다. 즉시 적용이나 관측 제거는 정당성을 검증하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/docs/terraform/basic-commands)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-19.md) · [다음 장 →](chapter-21.md)
