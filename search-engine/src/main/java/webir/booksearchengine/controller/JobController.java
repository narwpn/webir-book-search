package webir.booksearchengine.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import webir.booksearchengine.service.IndexService;

import org.jobrunr.scheduling.BackgroundJob;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;

@RestController
@RequestMapping("/jobs")
public class JobController {
    private final IndexService indexService;

    public JobController(IndexService indexService) {
        this.indexService = indexService;
    }

    @PostMapping("/indexAll")
    public ResponseEntity<String> scheduleIndexAllJob() {
        BackgroundJob.enqueue(() -> indexService.indexAll());
        return ResponseEntity.ok("Indexing job has been scheduled.");
    }

}
