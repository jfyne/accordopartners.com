import sbt._
import Keys._
import PlayProject._

object ApplicationBuild extends Build {

    val appName         = "accordopartners.com"
    val appVersion      = "1.0-SNAPSHOT"

    val appDependencies = Seq(
    )

    val main = PlayProject(appName, appVersion, appDependencies, mainLang = SCALA).settings(
    )

}
