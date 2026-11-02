# 선택 실습: GKE와 Kubernetes

유료 실습이다. 기본 개념 학습만 할 때는 manifest를 읽고 상태 진단을 설명한다. 실제 실행은 [준비](../00-getting-started/cost-safety.md) 후 수행한다. compute/container API, cluster 생성·kubectl 접근 권한, quota가 필요하다. 이미지 pull에 필요한 노드 identity의 Artifact Registry 접근과 네트워크도 확인한다.

```bash
gcloud services enable container.googleapis.com --project="$PROJECT_ID"
gcloud container clusters create-auto ace-lab-gke --region="$REGION" --project="$PROJECT_ID"
gcloud container clusters get-credentials ace-lab-gke --region="$REGION" --project="$PROJECT_ID"
kubectl config current-context
kubectl apply -f labs/kubernetes/deployment.yaml
kubectl rollout status deployment/ace-web -n ace-lab --timeout=300s
kubectl get pods,services -n ace-lab
kubectl logs -n ace-lab deployment/ace-web
kubectl port-forward -n ace-lab service/ace-web 8080:80
```

별도 터미널에서 `curl http://localhost:8080`을 실행한다. port-forward는 Ctrl+C로 종료한다. context가 다른 cluster라면 apply/delete를 실행하지 않는다. 이미지 pull 오류는 registry 존재·권한·egress를 확인한다. HPA는 `kubectl autoscale deployment ace-web -n ace-lab --cpu-percent=60 --min=1 --max=2`로 선택 확장한다. requests가 선언되어 있는지 확인한다.

## 정리

```bash
kubectl delete namespace ace-lab
gcloud container clusters delete ace-lab-gke --region="$REGION" --project="$PROJECT_ID"
gcloud compute disks list --project="$PROJECT_ID"
gcloud compute forwarding-rules list --project="$PROJECT_ID"
```

[GKE quickstart](https://cloud.google.com/kubernetes-engine/docs/quickstarts/create-cluster)
