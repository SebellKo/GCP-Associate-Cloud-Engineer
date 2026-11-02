# Chapter 21. Cloud Logging·Audit Logs·로그 라우팅

[← 이전 장](chapter-20.md) · [전체 목차](../README.md) · [다음 장 →](chapter-22.md)

공식 가이드 연결: **3.4** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 로그로 사건을 재구성한다
Cloud Logging은 이벤트 기록을 검색하고 Cloud Monitoring은 시계열 지표를 관찰한다. 로그를 metric으로 만들 수도 있지만 목적은 구분한다. resource type, timestamp, severity, trace/request ID로 범위를 좁힌다. 시간대가 다른 로그를 비교할 때 timestamp를 맞춘다.

Log Router의 sink는 필터에 맞는 로그를 지원 목적지로 라우팅한다. 생성 이전의 모든 로그를 자동 소급 복사하는 기능으로 보지 않는다. 목적지에 sink writer identity가 쓸 권한이 있어야 한다. 조직·폴더의 aggregated sink는 여러 프로젝트 로그를 모을 때 검토한다. log bucket 보관 기간과 제외 정책은 비용·감사 요구와 함께 결정한다.

Audit Logs는 누가 무엇을 했는지 조사한다. Admin Activity와 Data Access는 다르며 Data Access는 서비스별 기본 활성화·조회 권한·과금을 확인한다. VPC Flow Logs와 firewall rule logging은 네트워크 증거다. 모든 패킷의 전체 payload가 반드시 저장된다고 가정하지 않는다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. sink 생성이 과거 모든 로그를 옮기는가?

2. sink writer 권한이 필요한 이유는?

3. Admin Activity와 앱 오류 로그의 차이는?

## 3. 실습과 관찰

`gcloud logging read 'severity>=ERROR' --project="$PROJECT_ID" --limit=10`으로 최근 오류를 조회한다. Logs Explorer에서 resource.type 필터를 추가하고 시간 범위를 변경한다. 선택: 임시 sink를 만들고 writer identity와 목적지 권한을 확인한다. 끝나면 sink·목적지·보관 데이터를 삭제한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

세 프로젝트의 감사 로그를 중앙 프로젝트로 모으려 한다.

A. 범위에 맞는 aggregated sink와 목적지 권한을 구성한다
B. DNS 이름을 같게 만든다
C. 버킷 등급만 바꾼다
D. 각 VM CPU만 늘린다

### 문제 2 — 단일 선택

sink를 만들었지만 대상에 새 로그가 없다.

A. 조직 이름만 바꾼다
B. 필터·신규 로그 발생·writer identity의 대상 쓰기 권한을 확인한다
C. 모든 IAM을 삭제한다
D. VM을 무조건 재생성한다

### 문제 3 — 단일 선택

누가 방화벽 규칙을 바꿨는지 조사한다.

A. CPU metric만 본다
B. 관련 Admin Activity 감사 로그를 조회한다
C. 이미지 태그만 본다
D. Cloud Storage class만 본다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. 목적지에 기록해야 하기 때문이다.

3. 관리 작업 감사와 앱 실행 사건이다.

### 시험형 예제

**1. A** — 중앙 로그 라우팅과 쓰기 권한이 필요하다. CPU·DNS·등급은 로그를 모으지 않는다.

**2. B** — 라우팅 경로와 권한을 검증한다. IAM 전체 삭제나 VM 재생성·이름 변경은 무관한 파괴적 조치다.

**3. B** — 관리 행위와 주체는 감사 로그가 적합하다. 지표·태그·등급은 변경 주체 증거가 아니다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/logging/docs/routing/overview)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-20.md) · [다음 장 →](chapter-22.md)
