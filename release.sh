#!/usr/bin/env bash

#1. configuratin
APP_NAME="note-taker"
VERSION=${1:-"v1.0.0"}
BUILD_DIR="build"
REGISTRY="ghcr.io/thissidemayur"
IMAGE_NAME="$REGISTRY/$APP_NAME:$VERSION"
MAIN_GO_FILE_PATH="./cmd/note-taker/main.go"
#2. cleanup
echo "🧹 cleaning old build..."
rm -rf "$BUILD_DIR"
mkdir -p "$BUILD_DIR"

#3. Build binaries for multiple os & architecture
echo "⚙️  Building binaries..."
PLATFORMS=("linux/amd64" "darwin/amd64" "windows/amd64")

for PLATFORM in "${PLATFORMS[@]}"; do
	IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
	OUTPUT="$BUILD_DIR/${APP_NAME}-${GOOS}-${GOARCH}"
	[ "$GOOS" = "windows" ] && output+=".exe"

	echo "▶ Building for $GOOS/$GOARCH"
	GOOS="$GOOS" GOARCH=$GOARCH go build -o "$OUTPUT" $MAIN_GO_FILE_PATH

	tar -czf "$OUTPUT.tar.gz" -C "$BUILD_DIR" $(basename "$OUTPUT")

done

#4. Build docker image (optional)
echo "🐳 Building docker image... "
podman build -t $IMAGE_NAME .

#5. Push docker image (login require)
read -p "Do you want to push the image to GHCR/DockerHub? (y/n): "  PUSH 
if [[ $PUSH == "y" ]]; then
	  echo "🚀 Pushing image to registry..."
	  podman push $IMAGE_NAME
else 
	  echo "⏩ Skipped pushing image."
fi

#6 git tag and release 
read -p "Do you want to create a GitHub release for $VERSION? (y/n): " RELEASE
if [[ "$RELEASE" == "y" ]]; then
	  echo "🏷️ Tagging and creating release..."
	  git tag -a $VERSION -m "Release $VERSION"
	  git push origin $VERSION

	  echo "📦 Uploading binaries to GitHub release..."
	  gh release create "$VERSION" \
		$BUILD_DIR/*.tar.gz \
		--notes "Automated release for $APP_NAME $VERSION"
else
	    echo "⏩ Skipped GitHub release."
fi

echo "✅ Release $VERSION completed successfully!"
