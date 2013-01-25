/**
 * Fetch content from the google doc
 *
 * @author Josh Fyne
 */

package models

import play.api.libs.ws.WS
import play.api.Play
import play.api.Play.{configuration,current}
import scala.io._

object Content {

    /**
     * Content of the site
     *
     */
    var content:scala.collection.mutable.Map[String,String] = scala.collection.mutable.Map.empty[String,String]


    /**
     * Get a copy of the document
     *
     */
    def getContent = {
        WS.url(Play.configuration.getString("application.content").get).get().map { response =>
            response.body.split("\n").map(line => {
                val pair = line.split(",")
                content(pair(0)) = pair(1)
            })
        }
    }

    /**
     * Get a value from a key
     *
     */
    def fetch(key:String):Option[String] = {
        if (content.values.size == 0) getContent
        content.get(key)
    }
}
