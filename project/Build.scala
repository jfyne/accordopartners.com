import sbt._
import Keys._
import PlayProject._

object ApplicationBuild extends Build {

    val appName         = "accordopartners.com"
    val appVersion      = "1.0-SNAPSHOT"

    val appDependencies = Seq(
      "com.google.api-client" % "google-api-client" % "1.13.2-beta"
    )

    def customLessEntryPoints(base: File): PathFinder = (
        (base / "app" / "assets" / "stylesheets" / "bootstrap" * "bootstrap.less") +++
        (base / "app" / "assets" / "stylesheets" / "bootstrap" * "responsive.less") +++ 
        (base / "app" / "assets" / "stylesheets" * "*.less")
    )

    val main = PlayProject(appName, appVersion, appDependencies, mainLang = SCALA).settings(
      lessEntryPoints <<= baseDirectory(customLessEntryPoints)
    )

}
