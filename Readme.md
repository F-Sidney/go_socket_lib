# A simple Android native socket library created by golang and gomobile

## Build command:

```bash
gomobile bind -target=android -o="nsockets.aar" -javapkg="com.nxf.nsockets"
```

After build, it will create two files:

* nsockets.aar
* nsockets-sources.jar

## Use library in Android project:

1. Create new Android project in Android Stuido,
2. Copy those two files "nsockets.aar","nsockets-sources.jar" to the app folder,
3. Change "settings.gradle" file content as this:

> add those setting to "dependencyResolutionManagement"."repositories":

```groovy
flatDir {
	dirs 'app'
}
```

> the final gradle file should look like this:

```groovy
dependencyResolutionManagement {
    repositoriesMode.set(RepositoriesMode.FAIL_ON_PROJECT_REPOS)
    repositories {
        google()
        mavenCentral()
        flatDir {
            dirs 'app'
        }
    }
}
```

4. In "app" folder, open the file "build.gradle", go to the "dependencies" segment, add follow setting to the end:

```groovy
implementation (name:'nsockets', ext:'aar')
```
