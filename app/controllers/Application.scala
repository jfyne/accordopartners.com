package controllers

import play.api._
import play.api.mvc._

import models._

object Application extends Controller {

    def index = Action {
        val title = Content.fetch("home.title").get
        val subtitle = Content.fetch("home.subtitle").get
        Ok(views.html.index(title, subtitle))
    }
}
