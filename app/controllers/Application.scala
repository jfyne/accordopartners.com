package controllers

import play.api._
import play.api.mvc._
import play.api.cache.Cached
import play.api.Play.current

import models._

object Application extends Controller {

    /**
     * Home
     *
     */
    def index = Cached("acc.index", 86400) {
        Action {
            Async {
                val response = Content.getContent
                response.map({ sheet =>
                    Ok(views.html.index(Content.parse(sheet)))
                })
            }
        }
    }

    /**
     * About
     *
     */
    def about = Cached("acc.about", 86400) {
        Action {
            Async {
                val response = Content.getContent
                response.map({ sheet =>
                    Ok(views.html.about(Content.parse(sheet)))
                })
            }
        }
    }

    /**
     * Solutions
     *
     */
    def solutions = Cached("acc.solutions", 86400) {
        Action {
            Async {
                val response = Content.getContent
                response.map({ sheet =>
                    Ok(views.html.solutions(Content.parse(sheet)))
                })
            }
        }
    }

    /**
     * Experience
     *
     */
    def experience = Cached("acc.experience", 86400) {
        Action {
            Async {
                val expos = Expos.getContent
                expos.map({ expoSheet =>
                    val allExpos = Expos.parse(expoSheet)
                    val categories = allExpos.groupBy(e => e.categoryName)
                    Async {
                        val response = Content.getContent
                        response.map({ sheet =>
                            Ok(views.html.experience(Content.parse(sheet), categories))
                        })
                    }
                })
            }
        }
    }

    /**
     * Contact
     *
     */
    def contact = Cached("acc.contact", 86400) {
        Action {
            Async {
                val response = Content.getContent
                response.map({ sheet =>
                    Ok(views.html.contact(Content.parse(sheet)))
                })
            }
        }
    }

    /**
     * Expo
     *
     */
    def expos(tag:String) = Cached("acc.expos" + tag, 43200) {
        Action {
            Async {
                val response = Expos.getContent
                response.map({ sheet =>
                    val allExpos = Expos.parse(sheet)
                    val categories = allExpos.groupBy(e => e.categoryName)
                    val filteredExpos = if (tag != "Upcoming") allExpos.filter(e => e.categorySlug == tag) else allExpos
                    val title = {
                        if (tag == "Upcoming") {
                            tag
                        } else if (filteredExpos.length > 0) {
                            filteredExpos(0).categoryName
                        } else {
                            "None found"
                        }
                    }
                    Ok(views.html.expos(title, categories, filteredExpos))
                })
            }
        }
    }
}
