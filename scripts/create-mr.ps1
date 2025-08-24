# PowerShell скрипт для автоматического создания ветки и MR
param(
    [string]$CommitMessage = "",
    [string]$TargetBranch = "dev"
)

# Получаем текущую ветку
$currentBranch = git branch --show-current

# Если мы в master или dev, создаем новую ветку
if ($currentBranch -eq "master" -or $currentBranch -eq "dev") {
    $timestamp = Get-Date -Format "yyyyMMdd_HHmmss"
    $branchName = "feature/ai-update-$timestamp"
    
    Write-Host "Создаем новую ветку: $branchName" -ForegroundColor Green
    git checkout -b $branchName
    
    # Добавляем все изменения
    git add .
    
    # Создаем коммит
    if ($CommitMessage -eq "") {
        $CommitMessage = "AI update $timestamp"
    }
    git commit -m $CommitMessage
    
    # Пушим новую ветку
    git push origin $branchName
    
    Write-Host "Ветка $branchName создана и запушена" -ForegroundColor Green
    Write-Host "Создайте MR в ветку $TargetBranch через веб-интерфейс" -ForegroundColor Yellow
} else {
    # Если уже в feature ветке, просто коммитим и пушим
    git add .
    
    if ($CommitMessage -eq "") {
        $CommitMessage = "Update $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')"
    }
    git commit -m $CommitMessage
    
    git push origin $currentBranch
    
    Write-Host "Изменения запушены в ветку $currentBranch" -ForegroundColor Green
    Write-Host "Создайте MR в ветку $TargetBranch через веб-интерфейс" -ForegroundColor Yellow
}
