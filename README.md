# ToDo App on Kubernetes

## 概要

このプロジェクトは、Go言語で実装されたバックエンドAPIと、Reactで実装されたフロントエンドを組み合わせたシンプルなToDoアプリケーションです。

Kubernetes (Minikube) 上で動作するようにコンテナ化されており、Ingressコントローラを使用して外部からアクセス可能です。

## 技術スタック

-   **バックエンド:** Go, Gin Web Framework
-   **フロントエンド:** React, TypeScript, Vite
-   **コンテナ化:** Docker
-   **オーケストレーション:** Kubernetes (Minikube)
-   **APIゲートウェイ:** Nginx Ingress Controller

## 前提条件

以下のツールがローカル環境にインストールされている必要があります。

-   [Git](https://git-scm.com/)
-   [Docker Desktop](https://www.docker.com/products/docker-desktop/) (またはDocker Engine)
-   [Minikube](https://minikube.sigs.k8s.io/docs/start/)
-   [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
-   [Go](https://golang.org/doc/install) (バックエンドのローカル開発用)
-   [Node.js](https://nodejs.org/en/download/) (フロントエンドのローカル開発用)

## Minikube上での構築と実行手順

### 1. Minikubeクラスタの起動

Apple Silicon (M1/M2/M3) 環境の場合、QEMUドライバを使用することを推奨します。

```bash
minikube start --driver=qemu
```

### 2. Ingressアドオンの有効化

MinikubeクラスタでIngressコントローラを有効にします。

```bash
minikube addons enable ingress
```

### 3. Dockerイメージのビルド

MinikubeのDockerデーモンを使用して、バックエンドとフロントエンドのDockerイメージをビルドします。

```bash
# バックエンドイメージのビルド
minikube image build -t todo-backend backend/

# フロントエンドイメージのビルド
minikube image build -t todo-frontend frontend/
```

### 4. Kubernetesリソースのデプロイ

Deployment, Service, IngressリソースをKubernetesクラスタにデプロイします。

```bash
kubectl apply -f k8s/
```

Podが正常に起動していることを確認します。

```bash
kubectl get pods
```

### 5. アプリケーションへのアクセス

Ingress経由でアプリケーションにアクセスが可能です。

MinikubeクラスタのIPアドレスを取得します。

```bash
minikube ip
```

取得したIPアドレス（例: `192.168.105.2`）にブラウザでアクセスして、ToDoアプリが正常に動作することを確認してください。

例: `http://192.168.105.2`

-   ページが表示され、ToDoの初期リスト（空のはず）が表示されます。
-   新しいToDoを追加すると、リストに表示されます。
-   ページをリロードしても、追加したToDoが保持されています。

## (オプション) ローカル開発環境での実行

### バックエンド

`backend`ディレクトリに移動し、Goアプリケーションを直接実行します。

```bash
cd backend
go run ./cmd/server
```

### フロントエンド

`frontend`ディレクトリに移動し、依存関係をインストールしてから開発サーバーを起動します。

```bash
cd frontend
npm install
npm run dev
```

**注意:** ローカル開発環境で実行する場合、フロントエンドは`http://localhost:5173`で動作し、バックエンドは`http://localhost:8080`で動作します。フロントエンドの`App.tsx`のAPI呼び出しURLを`http://localhost:8080/api/todos`に戻すか、Viteのプロキシ設定を利用する必要があります。現在のコードはIngress経由でのアクセスを想定しているため、相対パス`/api/todos`を使用しています。
