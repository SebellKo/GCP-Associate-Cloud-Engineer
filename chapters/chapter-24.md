# Chapter 24. IAM 역할·정책·최소 권한

[← 이전 장](chapter-23.md) · [전체 목차](../README.md) · [다음 장 →](chapter-25.md)

공식 가이드 연결: **4.1** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 누가 어떤 자원에서 무엇을 하는가
Principal은 사용자·그룹·서비스 계정·federated identity 등 접근 주체다. Role은 permission 묶음이며 binding은 principal과 role을 resource 범위에 연결한다. 기본 역할 Owner/Editor/Viewer보다 작업에 맞는 predefined role을 먼저 검토한다. 필요한 permission 조합이 맞지 않으면 custom role의 지원·유지보수 부담을 검토한다.

권한은 가장 작은 적절한 범위에 부여한다. 한 버킷 객체를 읽는 앱에 프로젝트 전체 Storage Admin을 줄 필요는 없다. 상속·그룹·IAM Conditions·deny·조직 정책 등 실제 유효 권한에 영향을 주는 설정을 함께 조사한다.

인증(authentication)은 주체 확인, 인가(authorization)는 작업 허용 판단이다. 401/403을 단순 공식으로만 외우지 말고 서비스별 오류와 요청 identity를 확인한다. Policy Troubleshooter 등으로 특정 permission 실패를 조사하고 광범위 역할을 임시 해결책으로 남기지 않는다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. role과 permission의 관계는?

2. 최소 권한의 두 축은?

3. 그룹 소속이 유효 권한에 영향을 주는가?

## 3. 실습과 관찰

테스트 버킷에 전용 서비스 계정의 objectViewer 역할을 부여하고 describe로 정책을 읽는다. 다른 버킷까지 권한이 퍼지는지 범위를 설명한다. 실습이 끝나면 테스트 binding을 제거하고 계정을 삭제한다. 조직·회사 계정의 실제 권한은 바꾸지 않는다. 상세 명령은 Storage/통합 실습을 참고한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

앱이 버킷 하나의 객체만 읽어야 한다.

A. 프로젝트 Owner를 준다
B. 해당 버킷 범위의 Storage Object Viewer를 부여한다
C. 모든 버킷을 공개한다
D. 조직 Editor를 준다

### 문제 2 — 단일 선택

직원이 서비스 목록만 조회해야 한다. 첫 역할 선택 원칙은?

A. 비밀키를 공유한다
B. 모든 기본 역할을 동시에 준다
C. 필요 permission을 가진 좁은 predefined role을 검토한다
D. 무조건 Editor를 준다

### 문제 3 — 단일 선택

정책에 allow가 있는데 작업이 거부된다.

A. 인스턴스 이름을 바꾼다
B. 실제 principal·resource·permission과 deny·조건·상위 제약을 조사한다
C. DNS만 삭제한다
D. allow가 있으면 오류라고 단정한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. role은 permission의 묶음이다.

2. 허용 작업과 적용 자원 범위다.

3. 그렇다.

### 시험형 예제

**1. B** — 필요 작업과 자원 범위를 최소화한다. Owner·공개화·조직 Editor는 불필요한 접근을 넓힌다.

**2. C** — 작업에 맞는 predefined role을 우선 검토한다. 광범위 역할·키 공유는 최소 권한 원칙에 맞지 않는다.

**3. B** — allow 하나만으로 유효 권한을 단정할 수 없다. 나머지는 접근 평가를 진단하지 않는다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/iam/docs/overview)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-23.md) · [다음 장 →](chapter-25.md)
