# Chapter 22. Monitoring·경보·Ops Agent·Prometheus

[← 이전 장](chapter-21.md) · [전체 목차](../README.md) · [다음 장 →](chapter-23.md)

공식 가이드 연결: **3.4** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### 관측할 질문부터 정한다
Metric은 CPU·요청 수·오류율·지연처럼 시간에 따른 값을 나타낸다. Alerting policy는 조건·평가 기간·알림 채널로 구성한다. 임계값을 한 번 넘은 순간과 지속적으로 넘는 상황을 구분하고 결측 데이터 처리도 검토한다. 경보가 있어도 notification channel이 잘못됐으면 담당자가 모른다.

호스트 기본 metric만으로 VM 내부 메모리 사용량이나 앱 로그가 모두 수집된다고 생각하지 않는다. Ops Agent를 설치·설정해 guest metric과 로그를 수집한다. Managed Service for Prometheus는 Prometheus 방식의 metric 수집·조회 경로를 지원한다. custom metric의 label cardinality는 비용과 성능에 영향을 준다.

Monitoring의 metrics scope는 여러 프로젝트를 관측하는 범위를 관리한다. 리소스를 다른 프로젝트로 이동시키는 기능은 아니다. SLI는 관측 지표, SLO는 목표, error budget은 허용 실패량이다. ACE에서는 전문 SRE 수학보다 상태 확인·경보·수집 오류 진단을 우선 익힌다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. CPU 기본 metric이 있으면 guest memory metric도 자동으로 항상 있는가?

2. HPA와 alerting policy는 같은가?

3. 높은 cardinality가 문제가 되는 이유는?

## 3. 실습과 관찰

콘솔에서 VM CPU 또는 Cloud Run 요청 지표를 보고 alert condition·기간·채널을 설계한다. 테스트 채널 전달을 확인한다. Ops Agent는 VM 실습을 확장하는 선택 단계로 guest memory 수집 전후를 비교한다. 테스트용 정책·채널은 지우고 수집 로그 보관 비용도 확인한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

VM 메모리 사용량을 Monitoring에서 보고 싶지만 기본 지표에 없다.

A. VM 이름을 바꾼다
B. bucket을 공개한다
C. Ops Agent의 guest metric 수집을 설정한다
D. DNS TTL을 낮춘다

### 문제 2 — 단일 선택

오류율 경보는 발생했지만 운영자가 알림을 못 받았다.

A. CPU를 무조건 늘린다
B. sink를 전부 지운다
C. notification channel 설정·검증·전달 상태를 확인한다
D. DB를 복원한다

### 문제 3 — 단일 선택

Prometheus를 사용하는 앱 metric을 관리형으로 수집하려 한다.

A. Cloud DNS에 metric을 쓴다
B. Archive bucket만 만든다
C. Managed Service for Prometheus를 검토한다
D. 이미지 tag에 모든 값을 넣는다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. 아니다.

2. 규모 조정과 알림은 별개다.

3. 시계열 수·수집 비용·조회 부담을 키운다.

### 시험형 예제

**1. C** — guest memory는 적절한 수집 에이전트가 필요하다. DNS·공개화·이름은 수집을 추가하지 않는다.

**2. C** — 경보 발생과 알림 전달은 별도 단계다. 자원 증설·로그 삭제·복원은 채널 문제를 진단하지 않는다.

**3. C** — Prometheus 호환 수집 요구에 맞는다. DNS·태그·버킷만으로 시계열 수집 경로를 대체하지 못한다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/monitoring/agent/ops-agent)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-21.md) · [다음 장 →](chapter-23.md)
