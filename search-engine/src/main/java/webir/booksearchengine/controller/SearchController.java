package webir.booksearchengine.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import webir.booksearchengine.dto.SearchRequest;
import webir.booksearchengine.dto.SearchResponse;
import webir.booksearchengine.service.SearchService;
import org.springframework.web.bind.annotation.GetMapping;

@RestController
public class SearchController {

    private SearchService searchService;

    public SearchController(SearchService searchService) {
        this.searchService = searchService;
    }

    @GetMapping("/ping")
    public ResponseEntity<String> ping() {
        return ResponseEntity.ok("It works");
    }

    @PostMapping("/search")
    public SearchResponse search(@RequestBody SearchRequest searchRequest) {
        return searchService.search(searchRequest);
    }
}
