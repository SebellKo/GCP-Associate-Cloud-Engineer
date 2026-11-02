# Chapter 19. GKE·Kubernetes 배포와 운영

[← 이전 장](chapter-18.md) · [전체 목차](../README.md) · [다음 장 →](chapter-20.md)

공식 가이드 연결: **2.1 / 3.1** · 권장 2~3시간(통합 실습은 4~6시간)

## 1. 개념

### Kubernetes의 선언적 모델
Pod는 컨테이너 실행 단위, Deployment는 원하는 replica와 rollout을 관리하고, Service는 바뀌는 Pod에 안정적인 접근점을 제공한다. StatefulSet은 안정적인 identity·스토리지가 필요한 워크로드를 지원한다. ConfigMap은 일반 설정, Secret은 비밀 데이터를 전달하지만 base64가 암호화라는 뜻은 아니다.

GKE Standard는 노드 풀 설정·운영을 더 직접 제어한다. Autopilot은 인프라 관리를 줄이지만 Pod resource requests와 워크로드 설정은 여전히 사용자가 고려한다. regional/private cluster는 각각 가용성 배치와 네트워크 접근 특성을 나타내며 같은 의미가 아니다.

HPA는 Pod 수, VPA는 CPU·메모리 requests 조정, cluster autoscaler는 노드 capacity를 조정한다. 같은 CPU 지표로 HPA/VPA를 무분별하게 조합하지 않는다. Pending은 requests·스케줄링·capacity, CrashLoopBackOff는 앱 로그·설정, ImagePullBackOff는 이미지 가져오기 경로를 먼저 확인한다.

## 2. 개념 확인 문제

답을 먼저 적고 해설을 연다. 각 질문에 본인의 예시를 하나 붙인다.

1. HPA/VPA/노드 확장의 차이는?

2. Autopilot이면 requests가 무관한가?

3. Service가 Deployment 자체를 다시 시작하는가?

## 3. 실습과 관찰

[GKE 선택 실습](../labs/gke.md)을 수행한다. 생성 비용을 허용하지 않으면 manifest를 읽고 Pod/Deployment/Service 관계를 설명하는 것으로 시작한다. 실제 배포 시 rollout status, describe, logs를 확인한다. LoadBalancer 타입 대신 port-forward를 사용하며 완료 후 cluster와 외부 IP·디스크를 확인한다.

## 4. 실전 시험형 예제

학습용 자체 제작 문제다. 실제 출제 문항이나 시험 덤프가 아니다.

### 문제 1 — 단일 선택

CPU 사용률 증가에 따라 Pod 개수를 늘려야 한다.

A. VPA만으로 replica를 늘린다
B. DNS 레코드를 늘린다
C. 스냅샷 스케줄을 설정한다
D. HPA를 설정한다

### 문제 2 — 단일 선택

Pod가 반복 종료되어 CrashLoopBackOff다. 우선 진단은?

A. kubectl logs와 describe로 종료 원인·설정을 확인한다
B. 결제 계정을 바꾼다
C. 모든 IAM을 Owner로 만든다
D. 노드를 무조건 모두 삭제한다

### 문제 3 — 단일 선택

Kubernetes API가 꼭 필요하지만 노드 운영 부담은 줄이고 싶다.

A. Cloud DNS로 Pod를 실행한다
B. GKE Autopilot을 검토한다
C. VM 하나에만 모든 기능을 고정한다
D. Cloud Storage만 배포한다

<details>
<summary>개념 확인 정답과 시험형 문제 해설</summary>

### 개념 확인

1. Pod 수/Pod requests/노드 capacity다.

2. 스케줄링·자원 할당 등에 중요하다.

3. 안정된 접근점이며 rollout 제어와 다르다.

### 시험형 예제

**1. D** — HPA가 수평 확장이다. VPA는 requests 조정이고 백업·DNS는 Pod 수를 조절하지 않는다.

**2. A** — 앱 오류 증거를 먼저 조사한다. 무조건 삭제·결제·권한 확대는 진단을 생략한다.

**3. B** — Kubernetes와 관리 부담 요구를 함께 충족하는 선택이다. DNS·Storage는 실행 플랫폼이 아니며 단일 VM 고정은 클러스터 요구와 다르다.

</details>

## 5. 완료 체크리스트

- [ ] 핵심 개념을 문서 없이 설명하고 내 서비스의 예를 들었다.
- [ ] 개념 확인 3문항과 시험형 3문항을 풀고 오답 이유를 기록했다.
- [ ] 실습 결과 또는 설계 산출물을 남겼다.
- [ ] 생성 자원의 삭제와 잔존 비용을 확인했다(조회·설계만 했다면 해당 없음).

## 공식 자료

- [관련 제품/학습 문서](https://cloud.google.com/kubernetes-engine/docs/concepts/autopilot-overview)
- [시험 범위와 최신성 관리](../00-getting-started/exam-overview.md)

[← 이전 장](chapter-18.md) · [다음 장 →](chapter-20.md)
